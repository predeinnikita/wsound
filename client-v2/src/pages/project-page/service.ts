import axios from "axios";
import { AddAudioPayload, AudioList, Project } from "../../typing";

export const getProjectInfo = (projectId: number) => {
  return axios.get<Project>(`/api/projects/${projectId}`);
};

export const getProjectAudios = (projectId: number) => {
  return axios.get<AudioList>(`/api/audio?projectId=${projectId}`);
};

export const uploadAudioToStorage = (file: any) => {
  const fd = new FormData();
  fd.append("file", file);
  return fetch("/api/file-storage", {
    method: "POST",
    body: fd,
  });
};

export const addAudio = (payload: AddAudioPayload) => {
  return axios.post("/api/audio", payload);
};
