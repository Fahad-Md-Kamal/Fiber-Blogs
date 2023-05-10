import { Link } from "react-router-dom";

function Navigation() {
  return (
    <nav className="flex items-center justify-between flex-wrap bg-gray-800 p-6">
      <div className="flex items-center flex-shrink-0 text-white mr-6">
        <Link to="/" className="font-semibold text-xl tracking-tight">
          Tech Talk
        </Link>
      </div>
      <div className="block lg:hidden">
        <button
          className="flex items-center px-3 py-2 border rounded text-gray-300 border-gray-400 hover:text-white hover:border-white"
          aria-label="Toggle menu"
        >
          <svg
            className="fill-current h-3 w-3"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <title>Menu</title>
            <path
              d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"
              fillRule="evenodd"
            />
          </svg>
        </button>
      </div>
      <div className="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
        <div className="text-sm lg:flex-grow">
          <Link to="/" className="block mt-4 lg:inline-block lg:mt-0 text-gray-300 hover:text-white mr-4">Home</Link>
          <Link to="/about" className="block mt-4 lg:inline-block lg:mt-0 text-gray-300 hover:text-white mr-4">About</Link>
          <Link to="/auth/login" className="block mt-4 lg:inline-block lg:mt-0 text-gray-300 hover:text-white mr-4">Login</Link>
          <Link to="/auth/signup" className="block mt-4 lg:inline-block lg:mt-0 text-gray-300 hover:text-white mr-4">Signup</Link>
        </div>
      </div>
    </nav>
  );
}

export default Navigation;
