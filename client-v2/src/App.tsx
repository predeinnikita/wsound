import { BrowserRouter, Route, Routes, useNavigate } from "react-router";
import styles from "./app.module.css";
import { MainPage } from "./pages/main-page/MainPage";
import { ProjectPage } from "./pages/project-page/ProjectPage";
import { CreateProject } from "./pages/create-project-page/CreateProject";
import { Button, Layout, theme } from "antd";
import { SoundOutlined } from "@ant-design/icons";

const { Header, Content, Footer } = Layout;

const HeaderComponent = () => {
  const navigate = useNavigate();
  return (
    <Header>
      <SoundOutlined
        style={{ fontSize: "24px", color: "white", marginRight: "8px" }}
        onClick={() => navigate("/")}
      />
      <Button onClick={() => navigate("create-project")}>Create project</Button>
    </Header>
  );
};

function App() {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  return (
    <div className={styles.main}>
      <BrowserRouter>
        <Layout style={{ height: "100%", minHeight: "100vh", display: "flex" }}>
          <HeaderComponent />
          <Content
            style={{
              padding: "0 48px",
              marginTop: "24px",
              height: "100%",
              flex: 1,
            }}
          >
            <div
              style={{
                background: colorBgContainer,
                minHeight: 280,
                padding: 24,
                borderRadius: borderRadiusLG,
                flex: 1,
                height: "100%",
                display: "flex",
              }}
            >
              <Routes>
                <Route index element={<MainPage />} />
                <Route path="create-project" element={<CreateProject />} />
                <Route path="projects/:id" index element={<ProjectPage />} />
              </Routes>
            </div>
          </Content>
          <Footer style={{ textAlign: "center" }}>
            Ant Design ©{new Date().getFullYear()} Created by Ant UED
          </Footer>
        </Layout>
      </BrowserRouter>
    </div>
  );
}

export default App;
