import { FC, useCallback, useEffect, useMemo, useState } from "react";
import styles from "./ProjectPage.module.css";
import {
  Button,
  List,
  Skeleton,
  Space,
  Typography,
  Upload,
  UploadFile,
  UploadProps,
} from "antd";
import { useLocation } from "react-router";
import { addAudio, getProjectAudios, getProjectInfo } from "./service";
import { AudioList, Project } from "../../typing";
import { UploadOutlined } from "@ant-design/icons";

export const ProjectPage: FC = () => {
  const [currentProject, setCurrentProject] = useState<Project | null>(null);
  const [projectAudios, setProjectAudios] = useState<AudioList | null>(null);
  const [uploadFiles, setUploadFiles] = useState<UploadFile<any>[]>([]);
  const [isLoadingAudios, setIsLoadingAudios] = useState<boolean>(false);
  const [isLoadingAddAudios, setIsLoadingAddAudios] = useState<boolean>(false);

  const { pathname } = useLocation();

  const props: UploadProps = useMemo(() => {
    return {
      // beforeUpload: () => false,
      action: "/api/file-storage",
      onChange({ file, fileList }) {
        if (file.status !== "uploading") {
          console.log(file, fileList);
          setUploadFiles(fileList);
        }
      },
    };
  }, []);

  const getAudios = useCallback(() => {
    const projectId = pathname.split("/").at(-1);

    setIsLoadingAudios(true);
    getProjectAudios(+projectId!)
      .then(({ data }) => setProjectAudios(data))
      .finally(() => setIsLoadingAudios(false));
  }, [pathname]);

  const addAudioToProject = useCallback(async () => {
    const projectId = +pathname.split("/").at(-1)!;
    setIsLoadingAddAudios(true);
    for (let file of uploadFiles) {
      // const a = await uploadAudioToStorage(file);
      const audioId = file.response.id;
      await addAudio({
        name: file.name,
        project_id: projectId,
        storage_id: audioId,
      });
    }
    setIsLoadingAddAudios(false);
    getAudios();
    setUploadFiles([]);
  }, [getAudios, pathname, uploadFiles]);

  useEffect(() => {
    const projectId = pathname.split("/").at(-1);

    getProjectInfo(+projectId!).then(({ data }) => setCurrentProject(data));

    getAudios();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (!currentProject) {
    return <Skeleton />;
  }

  return (
    <div className={styles.main}>
      <Typography.Title>{currentProject.name}</Typography.Title>
      <Typography.Text>{currentProject.description}</Typography.Text>
      <Typography.Title level={3}>Audio</Typography.Title>
      <Space direction="vertical" style={{ width: "100%" }}>
        <List
          loading={isLoadingAudios}
          size="large"
          bordered
          dataSource={projectAudios?.audios || []}
          renderItem={({ name }) => (
            <List.Item>
              <Typography.Text>{name}</Typography.Text>
            </List.Item>
          )}
        />
        <Upload {...props}>
          <Button icon={<UploadOutlined />}>Upload</Button>
        </Upload>
        <Button loading={isLoadingAddAudios} onClick={addAudioToProject}>
          Add to project
        </Button>
      </Space>
    </div>
  );
};
