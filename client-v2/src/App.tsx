import { BrowserRouter, Route, Routes, useNavigate } from "react-router";
import styles from "./app.module.scss";
import { MainPage } from "./pages/main-page/MainPage";
import { ProjectPage } from "./pages/project-page/ProjectPage";
import { CreateProject } from "./pages/create-project-page/CreateProject";
import { ConfigProvider, Layout, theme } from "antd";

const { Header, Content, Footer } = Layout;

const HeaderComponent = () => {
  const navigate = useNavigate();
  return (
    <Header style={{ background: "#302E2E"}} onClick={() => navigate("/")}>
        <img src="/logo.svg" alt="logo" />
        <span>WSound</span>
    </Header>
  );
};

function App() {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  return (
      <ConfigProvider
          theme={{
              token: {
                  colorPrimary: '#302E2E',
              },
          }}
      >
          <div className={styles.main}>
              <BrowserRouter>
                  <Layout style={{height: "100%", minHeight: "100vh", display: "flex"}}>
                      <HeaderComponent/>
                      <main>
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
                                  <Route index element={<MainPage/>}/>
                                  <Route path="create-project" element={<CreateProject/>}/>
                                  <Route path="projects/:id" index element={<ProjectPage/>}/>
                              </Routes>
                          </div>
                      </main>
                      <Footer style={{textAlign: "center"}}>
                          Ant Design ©{new Date().getFullYear()} Created by Ant UED
                      </Footer>
                  </Layout>
              </BrowserRouter>
          </div>
      </ConfigProvider>
  );
}

export default App;
