import React from "react";
import { Outlet } from "react-router";

type Props = {};

const AuthenticationLayout: React.FC = (props: Props) => {
  return (
    <>
      <h2 className="text-center text-3xl font-extrabold text-gray-900 mt-20">
        Hi! <br /> Welcome to Tech Talk
      </h2>
      <div className="max-h-screen flex flex-col justify-center py-3 sm:px-6 lg:px-8">
          <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
            <Outlet />
          </div>
      </div>
    </>
  );
};

export default AuthenticationLayout;
