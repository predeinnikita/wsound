const fileInput = document.querySelector(".audio_input");
const sendButton = document.querySelector(".send_button");
const resultCont = document.querySelector(".result");

fileInput.addEventListener("change", (e) => {
  console.log(fileInput.files);
});

sendButton.addEventListener("click", () => {
  const fd = new FormData();
  fd.append("audioFile", fileInput.files[0]);

  resultCont.textContent = "loading...";

  fetch("/api/v1/audio/info", {
    method: "POST",
    body: fd,
  })
    .then((res) => res.json())
    .then((res) => {
      resultCont.textContent = Object.entries(res)
        .map(([key, value]) => `${key}: ${value}`)
        .join("\n\n");
    })
    .catch((e) => (resultCont.textContent = "error: " + e));
});
