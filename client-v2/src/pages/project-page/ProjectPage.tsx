import { FC, useCallback, useEffect, useMemo, useState } from "react";
import styles from "./ProjectPage.module.css";
import {
  Breadcrumb,
  Button,
  Flex,
  List,
  Skeleton,
  Space,
  Tag,
  Typography,
  Upload,
  UploadProps,
} from "antd";
import { useLocation, useNavigate } from "react-router";
import {
  addAudio,
  deleteAudio,
  getProjectAudios,
  getProjectInfo,
} from "./service";
import { AudioList, Project } from "../../typing";
import { DeleteOutlined, UploadOutlined } from "@ant-design/icons";

export const ProjectPage: FC = () => {
  const [currentProject, setCurrentProject] = useState<Project | null>(null);
  const [projectAudios, setProjectAudios] = useState<AudioList | null>(null);
  const [isLoadingAudios, setIsLoadingAudios] = useState<boolean>(false);
  const [isLoadingAddAudios, setIsLoadingAddAudios] = useState<boolean>(false);
  const [isLoadingDeleteAudios, setIsLoadingDeleteAudios] = useState<number>(0);
  const [isUploading, setIsUploading] = useState<boolean>(false);

  const { pathname } = useLocation();

  const navigate = useNavigate();

  const getAudios = useCallback(() => {
    const projectId = pathname.split("/").at(-1);

    setIsLoadingAudios(true);
    getProjectAudios(+projectId!)
      .then(({ data }) => setProjectAudios(data))
      .finally(() => setIsLoadingAudios(false));
  }, [pathname]);

  const props: UploadProps = useMemo(() => {
    return {
      action: "/api/file-storage",
      onChange({ file, fileList, event }) {
        if (typeof event?.percent === "number" && event.percent !== 100) {
          setIsUploading(true);
        } else {
          setIsUploading(false);
        }

        if (file.status !== "uploading" && file.status !== "removed") {
          console.log(file, fileList);

          const projectId = +pathname.split("/").at(-1)!;
          setIsLoadingAddAudios(true);
          addAudio({
            name: file.name,
            project_id: projectId,
            storage_id: file.response.id,
          }).finally(() => {
            setIsLoadingAddAudios(false);
            getAudios();
          });
        }
      },
      itemRender(_, file, __, { remove }) {
        if (file.status === "done") {
          remove();
        }
        return null;
      },
    };
  }, [getAudios, pathname]);

  const onClickDeleteAudio = useCallback(
    async (audioId: number) => {
      setIsLoadingDeleteAudios(audioId);
      await deleteAudio(audioId);
      setIsLoadingDeleteAudios(0);

      getAudios();
    },
    [getAudios]
  );

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
    <Space className={styles.main} direction="vertical">
      <Breadcrumb
        items={[
          {
            // eslint-disable-next-line jsx-a11y/anchor-is-valid
            title: <a onClick={() => navigate("/")}>Project list</a>,
          },
          {
            title: currentProject.name,
          },
        ]}
      />
      <Typography.Title>{currentProject.name}</Typography.Title>
      <Typography.Text>{currentProject.description}</Typography.Text>
      <Flex align="center" justify="space-between">
        <Typography.Title level={3} style={{ margin: "12px 0" }}>
          Audio
        </Typography.Title>
        <Upload {...props}>
          <Button
            type="primary"
            icon={<UploadOutlined />}
            loading={isUploading}
          >
            Upload
          </Button>
        </Upload>
      </Flex>
      <Space direction="vertical" style={{ width: "100%" }}>
        {/* <Dragger> */}
        <List
          loading={isLoadingAudios || isLoadingAddAudios}
          size="large"
          bordered
          dataSource={projectAudios?.audios || []}
          renderItem={({ name, status, id }) => (
            <List.Item>
              <Typography.Text>{name}</Typography.Text>
              <div style={{ display: "flex", alignItems: "center" }}>
                <Tag>{status}</Tag>
                <Button
                  danger
                  shape="circle"
                  icon={<DeleteOutlined />}
                  onClick={() => onClickDeleteAudio(id)}
                  loading={isLoadingDeleteAudios === id}
                />
              </div>
            </List.Item>
          )}
        />
        {/* </Dragger> */}
      </Space>
    </Space>
  );
};
