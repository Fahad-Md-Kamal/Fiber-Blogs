import { Route, Routes } from "react-router-dom";
import Navigation from "./parts/Navigation";
import About from "./pages/about/About";
import Home from "./pages/home/Home";
import Login from "./pages/authentications/Login";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

function App() {
  return (
    <>
      <Navigation />
      <ToastContainer />
      <div className="flex">
        <div className="flex-1 p-4">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/login" element={<Login />} />
          </Routes>
        </div>
      </div>
    </>
  );
}

export default App;
