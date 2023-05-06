import { Route, Routes } from "react-router-dom";
import Navigation from "./parts/Navigation";
import About from "./pages/about/About";
import Home from "./pages/home/Home";

function App() {
  return (
    <>
      <Navigation />
      <div className="flex">
        <div className="flex-1 p-4">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
          </Routes>
        </div>
      </div>
    </>
  );
}

export default App
