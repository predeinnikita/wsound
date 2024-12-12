import axios from "axios";
import { AddAudioPayload, AudioList, Project } from "../../typing";
import { CreateProjectForm } from "../create-project-page/service";

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

export const deleteAudio = (audioId: number) => {
  return axios.delete(`/api/audio/${audioId}`);
};

export const deleteProject = (projectId: number) => {
  return axios.delete(`/api/projects/${projectId}`);
};

export const editProject = (projectId: number, payload: CreateProjectForm) => {
  return axios.patch(`/api/projects/${projectId}`, payload);
};

export const editAudio = (audioId: number, payload: { name: string }) => {
  return axios.patch(`/api/audio/${audioId}`, payload);
};
