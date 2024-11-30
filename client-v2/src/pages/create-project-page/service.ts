import axios from "axios";
import { Project } from "../../typing";

export type CreateProjectForm = {
  name: string;
  description: string;
};

export const createProject = (payload: CreateProjectForm) => {
  return axios.post<Project>("/api/projects", payload);
};
