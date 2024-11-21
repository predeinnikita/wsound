const fileInput = document.querySelector("#addAudio");
const sendButton = document.querySelector("#addAudioBtn");

fileInput.addEventListener("change", (e) => {
  console.log(fileInput.files);
});
sendButton.addEventListener("click", () => {
  const fd = new FormData();
  fd.append("audioFile", fileInput.files[0]);
  fetch("/api/audio", {
    method: "POST",
    body: fd,
  })
    .then((res) => res.json())
    // .then((res) => {
    //   resultCont.textContent = Object.entries(res)
    //     .map(([key, value]) => `${key}: ${value}`)
    //     .join("\n\n");
    // })
    .catch((e) => alert("error: " + e));
});

const getProjectInfo = async () => {
  const title = document.getElementById("projectTitle");
  const description = document.getElementById("projectDescription");
  const audioList = document.getElementById("audioList");

  const projId = window.location.pathname.split("/").at(-1);

  const data = await (await fetch("/api/projects/" + projId)).json();
  const { audios } = await (
    await fetch("/api/audio?projectId=" + projId)
  ).json();

  title.textContent = data.name;
  description.textContent = data.description;

  audios.forEach((audio) => {
    const audioItem = document.createElement("a");
    audioItem.href = `audio-page.html?id=${audio.id}`;
    audioItem.className =
      "list-group-item list-group-item-action d-flex justify-content-between align-items-center";
    audioItem.innerHTML = `
            ${audio.name}
            <span class="badge badge-info">${audio.label}</span>
        `;
    audioList.appendChild(audioItem);
  });
};

getProjectInfo();

function deleteProject() {
  alert("Project deleted (mock action)");
}
