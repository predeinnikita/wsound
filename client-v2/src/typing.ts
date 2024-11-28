export type Project = {
  id: number;
  name: string;
  description: string;
  created_at: string;
};

export type Audio = {
  id: number;
  storage_id: string;
  name: string;
  project_id: number;
  created_at: string;
};

export type AudioList = {
  audios: Audio[];
  total: number;
};

export type AddAudioPayload = Pick<Audio, "name" | "project_id" | "storage_id">;
