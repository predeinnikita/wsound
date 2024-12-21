import { FC, useCallback, useEffect, useMemo, useState } from "react";
import styles from "./ProjectPage.module.scss";
import {
  Breadcrumb,
  Button,
  Flex,
  Form,
  Input,
  List,
  notification,
  Skeleton,
  Space,
  Tag,
  Tooltip,
  Typography,
  Upload,
  UploadProps,
} from "antd";
import { useLocation, useNavigate } from "react-router";
import {
  addAudio,
  deleteAudio,
  deleteProject,
  editAudio,
  editProject,
  getProjectAudios,
  getProjectInfo,
} from "./service";
import { AudioList, Project } from "../../typing";
import {
  DeleteOutlined,
  DownloadOutlined,
  EditOutlined,
  ExportOutlined, FileZipOutlined,
  SaveOutlined,
  UploadOutlined,
} from "@ant-design/icons";
import { CreateProjectForm } from "../create-project-page/service";

const { Dragger } = Upload;

export const ProjectPage: FC = () => {
  const [currentProject, setCurrentProject] = useState<Project | null>(null);
  const [projectAudios, setProjectAudios] = useState<AudioList | null>(null);
  const [isEditMode, setIsEditMode] = useState<boolean>(false);
  const [currentEditAudio, setCurrentEditAudio] = useState<number>(0);
  const [isDrop, setIsDrop] = useState<boolean>(false);

  const [isLoadingAudios, setIsLoadingAudios] = useState<boolean>(false);
  const [isLoadingAddAudios, setIsLoadingAddAudios] = useState<boolean>(false);
  const [isLoadingDeleteAudios, setIsLoadingDeleteAudios] = useState<number>(0);
  const [isUploading, setIsUploading] = useState<boolean>(false);
  const [isLoadingDeleteProject, setIsLoadingDeleteProject] =
    useState<boolean>(false);
  const [isLoadigSaveProjectChanges, setIsLoadingSaveProjectsChanges] =
    useState<boolean>(false);
  const [isLoadigSaveAudioChanges, setIsLoadingSaveAudioChanges] =
    useState<number>(0);

  const [form] = Form.useForm<CreateProjectForm>();
  const [formValid, setFormValid] = useState<boolean>(false);
  const formValues = Form.useWatch([], form);

  const [editAudioForm] = Form.useForm<{ name: string }>();
  const [editAudioFormValid, setEditAudioFormValid] = useState<boolean>(true);
  const editAudioFormValues = Form.useWatch([], editAudioForm);

  const [notificationApi, notificationContext] = notification.useNotification();

  const { pathname } = useLocation();

  const navigate = useNavigate();

  const projectId = pathname.split("/").at(-1);

  const onUploadError = useCallback(
    (message: string, description: string) => {
      notificationApi.error({
        message: message,
        description: description,
      });
    },
    [notificationApi]
  );

  const getAudios = useCallback(() => {
    setIsLoadingAudios(true);
    getProjectAudios(+projectId!)
      .then(({ data }) => setProjectAudios(data))
      .finally(() => setIsLoadingAudios(false));
  }, [projectId]);

  const getCurrentProject = useCallback(async () => {
    const projectId = pathname.split("/").at(-1);

    await getProjectInfo(+projectId!).then(({ data }) =>
      setCurrentProject(data)
    );
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

        if (file.status === "error") {
          onUploadError(file.response.message, file.response.detail);
          return;
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

        if (fileList.length === 0) {
          setIsDrop(false);
        }
      },
      itemRender(_, file, __, { remove }) {
        if (file.status === "done") {
          remove();
        }
        return null;
      },
    };
  }, [getAudios, onUploadError, pathname]);

  const onClickDeleteAudio = useCallback(
    async (audioId: number) => {
      setIsLoadingDeleteAudios(audioId);
      await deleteAudio(audioId);
      setIsLoadingDeleteAudios(0);

      getAudios();
    },
    [getAudios]
  );

  const onCickEditOrSaveButton = useCallback(async () => {
    if (isEditMode) {
      setIsLoadingSaveProjectsChanges(true);
      const projectId = pathname.split("/").at(-1) || 0;
      await editProject(+projectId, form.getFieldsValue());
      await getCurrentProject();

      setIsLoadingSaveProjectsChanges(false);
      setIsEditMode(false);
    } else {
      setIsEditMode(true);
    }
  }, [form, getCurrentProject, isEditMode, pathname]);

  const onClickDeleteProject = useCallback(async () => {
    setIsLoadingDeleteProject(true);

    const projectId = pathname.split("/").at(-1) || 0;

    await deleteProject(+projectId);

    setIsLoadingDeleteProject(false);
    navigate({ pathname: "/" }, { replace: true });
  }, [navigate, pathname]);

  const onClickEditOrSaveAudioChanges = useCallback(
    async (audioId: number) => {
      if (currentEditAudio === audioId) {
        //save
        setIsLoadingSaveAudioChanges(audioId);
        await editAudio(audioId, editAudioForm.getFieldsValue());

        setCurrentEditAudio(0);
        setIsLoadingSaveAudioChanges(0);

        getAudios();
      } else {
        setCurrentEditAudio(audioId);
      }
    },
    [currentEditAudio, editAudioForm, getAudios]
  );

  useEffect(() => {
    form
      .validateFields({ validateOnly: true })
      .then(() => setFormValid(true))
      .catch(() => setFormValid(false));
  }, [form, formValues]);

  useEffect(() => {
    editAudioForm
      .validateFields({ validateOnly: true })
      .then(() => setEditAudioFormValid(true))
      .catch(() => setEditAudioFormValid(false));
  }, [editAudioForm, editAudioFormValues]);

  useEffect(() => {
    getCurrentProject();

    getAudios();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    const dragopver = () => {
      setIsDrop((prev) => (!prev ? !prev : prev));
    };

    const dragleave = (e: any) => {
      if (e.target.tagName === "MAIN") {
        setIsDrop(false);
      }
    };

    window.addEventListener("dragover", dragopver);
    window.addEventListener("dragleave", dragleave);

    return () => {
      window.removeEventListener("dragover", dragopver);
      window.removeEventListener("dragleave", dragleave);
    };
  }, []);

  if (!currentProject) {
    return <Skeleton />;
  }

  return (
    <>
      {notificationContext}
      <Space className={styles.main} direction="vertical">
        <Flex align="center" justify="space-between">
          <Breadcrumb
            items={[
              {
                // eslint-disable-next-line jsx-a11y/anchor-is-valid
                title: <a onClick={() => navigate("/")}>Проекты</a>,
              },
              {
                title: currentProject.name,
              },
            ]}
          />
          <Flex gap={8}>
            {isEditMode && (
              <Button
                danger
                icon={<DeleteOutlined />}
                onClick={onClickDeleteProject}
                loading={isLoadingDeleteProject}
              />
            )}
            <Button
              type="dashed"
              disabled={!formValid}
              icon={isEditMode ? <SaveOutlined /> : <EditOutlined />}
              onClick={onCickEditOrSaveButton}
              loading={isLoadigSaveProjectChanges}
            />
          </Flex>
        </Flex>
        {isEditMode && (
          <Form
            layout="vertical"
            form={form}
            initialValues={{
              name: currentProject.name,
              description: currentProject.description,
            }}
          >
            <Form.Item
              label="Название"
              name="name"
              rules={[
                { required: true, message: "Пожалуйста, введите название" },
              ]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              label="Описание"
              name="description"
              rules={[
                { required: true, message: "Пожалуйста, введите описание" },
              ]}
            >
              <Input />
            </Form.Item>
          </Form>
        )}
        {!isEditMode && (
          <>
            <Typography.Title>{currentProject.name}</Typography.Title>
            <Typography.Text>{currentProject.description}</Typography.Text>
          </>
        )}
        <div className={styles.audioTitle}>
          <Typography.Title level={3} style={{ margin: "12px 0" }}>
            Список аудио
          </Typography.Title>
          <Flex gap="10px" justify="space-between" align="center">
            <Button
              type="default"
              href={`/api/export/excel/project/${projectId}`}
              icon={<ExportOutlined />}
            >
              Экспорт
            </Button>
            <Upload {...props} multiple>
              <Button
                type="primary"
                icon={<UploadOutlined />}
                loading={isUploading}
              >
                Загрузить
              </Button>
            </Upload>
          </Flex>
        </div>
        <Space direction="vertical" style={{ width: "100%" }}>
          {isDrop && <Dragger {...props} multiple height={300}></Dragger>}
          {!isDrop && (
            <List
              loading={isLoadingAudios || isLoadingAddAudios}
              size="large"
              bordered
              dataSource={projectAudios?.audios || []}
              renderItem={({ name, status, id, storage_id }) => (
                <List.Item>
                  <div className={styles.audio}>
                    <div>
                      {currentEditAudio !== id && (
                        <Typography.Text>{name}</Typography.Text>
                      )}
                      {currentEditAudio === id && (
                        <Form form={editAudioForm} initialValues={{ name }}>
                          <Form.Item
                            style={{ margin: 0 }}
                            name="name"
                            rules={[
                              {
                                required: true,
                                message: "Пожалуйста, введите название",
                              },
                            ]}
                          >
                            <Input />
                          </Form.Item>
                        </Form>
                      )}
                    </div>
                    <div style={{ display: "flex", alignItems: "center" }}>
                      <Tag>{status === "wolf" ? "Волк" : "Не волк"}</Tag>
                      <Tooltip placement="topRight" title="Редактировать аудио">
                        <Button
                            type="text"
                            disabled={!editAudioFormValid}
                            onClick={() => onClickEditOrSaveAudioChanges(id)}
                            icon={
                              currentEditAudio === id ? (
                                  <SaveOutlined />
                              ) : (
                                  <EditOutlined />
                              )
                            }
                            loading={isLoadigSaveAudioChanges === id}
                        />
                      </Tooltip>
                      <Tooltip placement="topRight" title="Скачать интервалы с воем">
                        <Button
                            disabled={status === 'not_wolf'}
                            type="text"
                            icon={<FileZipOutlined />}
                            href={`/api/export/zip/audio/${id}`}
                        />
                      </Tooltip>
                      <Tooltip placement="topRight" title="Экспорт Excel">
                        <Button
                            type="text"
                            icon={<ExportOutlined />}
                            href={`/api/export/excel/audio/${id}`}
                        />
                      </Tooltip>
                      <Tooltip placement="topRight" title="Скачать аудио">
                        <Button
                            type="text"
                            icon={<DownloadOutlined />}
                            href={`/api/file-storage/${storage_id}`}
                        />
                      </Tooltip>
                      <Tooltip placement="topRight" title="Удалить аудио">
                        <Button
                            danger
                            type="text"
                            icon={<DeleteOutlined />}
                            onClick={() => onClickDeleteAudio(id)}
                            loading={isLoadingDeleteAudios === id}
                        />
                      </Tooltip>
                    </div>
                  </div>
                </List.Item>
              )}
            />
          )}
        </Space>
      </Space>
    </>
  );
};
