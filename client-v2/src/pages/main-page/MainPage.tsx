import { FC, useEffect, useState } from "react";
import styles from "./MainPage.module.css";
import { Typography, List } from "antd";
import { Projects, getProjects } from "./service";
import { useNavigate } from "react-router";

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
      <Typography.Title>Projects</Typography.Title>
      <List
        loading={isLoading}
        size="large"
        bordered
        dataSource={projects?.projects || []}
        renderItem={({ name, id }) => (
          <List.Item onClick={() => navigate({ pathname: `projects/${id}` })}>
            <Typography.Text>{name}</Typography.Text>
          </List.Item>
        )}
      />
    </div>
  );
};
