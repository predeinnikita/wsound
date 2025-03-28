import { FC, useEffect, useState } from "react";
import styles from "./MainPage.module.scss";
import { Typography, List, Button, Flex } from "antd";
import { Projects, getProjects } from "./service";
import { useNavigate } from "react-router";
import { ExportOutlined } from "@ant-design/icons";

export const MainPage: FC = () => {
  const [projects, setProjects] = useState<Projects | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const navigate = useNavigate();

  useEffect(() => {
    setIsLoading(true);
    getProjects()
      .then(({ data }) => setProjects(data))
      .finally(() => setIsLoading(false));
  }, []);

  return (
    <div className={styles.main}>
      <div className={styles.title}>
        <Typography.Title>Проекты</Typography.Title>
          <Flex gap="10px" justify="space-between" align="center">
            <Button
                type="default"
                href="/api/export/excel"
                icon={<ExportOutlined />}
            >
                Экспорт
            </Button>
            <Button onClick={() => navigate("create-project")} type="primary">
              Создать проект
            </Button>
          </Flex>
      </div>
      <List
        loading={isLoading}
        size="large"
        bordered
        dataSource={projects?.projects || []}
        renderItem={({ name, id }) => (
          <List.Item className={styles.item} onClick={() => navigate({ pathname: `projects/${id}` })}>
            <Typography.Text>{name}</Typography.Text>
          </List.Item>
        )}
      />
    </div>
  );
};
