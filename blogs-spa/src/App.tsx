import { Route, Routes } from "react-router-dom";
import "react-toastify/dist/ReactToastify.css";
import About from "./pages/about/About";
import { ToastContainer } from "react-toastify";
import AuthenticationLayout from "./pages/layout/AuthenticationLayout";
import Home from "./pages/home/Home";
import AuthProvider from "./contexts/AuthProvider";

function App() {
  return (
    <>
      <ToastContainer />
      <AuthProvider>
        <Routes>
          <Route
            path="/"
            element={<Home />}
            children={[
              <Route path="/auth" element={<AuthenticationLayout />} />,
              <Route path="/about" element={<About />} />,
            ]}
          />
        </Routes>
      </AuthProvider>
    </>
  );
}

export default App;
