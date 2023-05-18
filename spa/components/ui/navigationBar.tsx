import { ApplicationRoutesEnum } from "@/lib/routesList";
import Link from "next/link";
import React from "react";

const Navbar = () => {
  return (
    <nav className="bg-gray-800">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-16">
          <div className="flex-shrink-0">
            <span className="text-white font-bold text-lg">Logo</span>
          </div>
          <div className="flex">
          <Link className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium" 
          href={ApplicationRoutesEnum.HOME}>Home</Link>
          <Link className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium" 
          href={ApplicationRoutesEnum.LOGIN}>Login</Link>
          <Link className="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium" 
          href={ApplicationRoutesEnum.SINGUP}>Signup</Link>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
