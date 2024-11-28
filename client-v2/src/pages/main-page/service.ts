import axios from "axios";
import { Project } from "../../typing";

export type Projects = {
  projects: Project[];
  total: number;
};

export const getProjects = () => {
  return axios.get<Projects>("/api/projects");
};
