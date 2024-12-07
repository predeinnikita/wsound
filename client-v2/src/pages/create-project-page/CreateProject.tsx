import { FC, useCallback, useState } from "react";
import styles from "./CreateProject.module.css";
import { Breadcrumb, Typography } from "antd";
import { Button, Form, Input } from "antd";
import { createProject, CreateProjectForm } from "./service";
import { useNavigate } from "react-router";

export const CreateProject: FC = () => {
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const [form] = Form.useForm<CreateProjectForm>();

  const navigate = useNavigate();

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
          rules={[{ required: true, message: "Please input name!" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          label="Описание"
          name="description"
          rules={[{ required: true, message: "Please input description!" }]}
        >
          <Input />
        </Form.Item>
        <Form.Item>
          <Button
            loading={isLoading}
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
