import { Route, Routes } from "react-router-dom";
import "react-toastify/dist/ReactToastify.css";
import Navigation from "./parts/Navigation";
import About from "./pages/about/About";
import Login from "./pages/authentications/Login";
import { ToastContainer } from "react-toastify";
import AuthenticationLayout from "./pages/layout/AuthenticationLayout";
import Signup from "./pages/authentications/Signup";
import Home from "./pages/home/Home";

function App() {
  return (
    <>
      <ToastContainer />
      <Navigation />

      <Routes>
        <Route path="/" element={<Home />} />
        <Route
          path="/auth"
          element={<AuthenticationLayout />}
          children={[
            <Route path="login" element={<Login />} />,
            <Route path="signup" element={<Signup />} />,
          ]}
        />
        <Route path="/about" element={<About />} />
      </Routes>
    </>
  );
}

export default App;
