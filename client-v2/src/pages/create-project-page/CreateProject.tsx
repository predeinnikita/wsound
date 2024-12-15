import { FC, useCallback, useEffect, useState } from "react";
import styles from "./CreateProject.module.css";
import { Breadcrumb, Typography } from "antd";
import { Button, Form, Input } from "antd";
import { createProject, CreateProjectForm } from "./service";
import { useNavigate } from "react-router";

export const CreateProject: FC = () => {
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const [form] = Form.useForm<CreateProjectForm>();

  const navigate = useNavigate();

  const [valid, setValid] = useState<boolean>(false);
  const formValues = Form.useWatch([], form);
  useEffect(() => {
    form
      .validateFields({ validateOnly: true })
      .then(() => setValid(true))
      .catch(() => setValid(false));
  }, [form, formValues]);

  const onSubmit = useCallback(() => {
    setIsLoading(true);
    createProject(form.getFieldsValue())
      .then(({ data: { id } }) => navigate(`/projects/${id}`))
      .finally(() => setIsLoading(false));
  }, [form, navigate]);

  return (
    <div className={styles.main}>
      <Breadcrumb
        items={[
          {
            // eslint-disable-next-line jsx-a11y/anchor-is-valid
            title: <a onClick={() => navigate("/")}>Проекты</a>,
          },
          {
            title: "Создание проекта",
          },
        ]}
      />
      <Typography.Title>Создать проект</Typography.Title>
      <Form
        layout="vertical"
        form={form}
        initialValues={{ name: "", description: "" }}
      >
        <Form.Item
          label="Название"
          name="name"
          rules={[{ required: true, message: "Пожалуйста, введите название" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Описание"
          name="description"
          rules={[{ required: true, message: "Пожалуйста, введите описание" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item>
          <Button
            loading={isLoading}
            disabled={!valid}
            type="primary"
            onClick={onSubmit}
            htmlType="submit"
          >
            Создать
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};
