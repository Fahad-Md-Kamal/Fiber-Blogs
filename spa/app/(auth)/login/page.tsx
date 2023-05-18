"use client";

import React, { useState } from "react";
import { FaEye, FaEyeSlash } from "react-icons/fa";

import { useForm, SubmitHandler } from "react-hook-form";
import Link from "next/link";
import { ApplicationRoutesEnum } from "@/lib/routesList";

interface IFormInput {
  username: String;
  password: String;
}

type Props = {};

const page = (props: Props) => {
  const { register, handleSubmit, formState } = useForm<IFormInput>({defaultValues: async () => {
    const response = await fetch("https://jsonplaceholder.typicode.com/posts/1");
    const data = await response.json();
    return {
      username: data.title,
      password: data.userId
    }
    
  }});
  const { errors } = formState;
  const onSubmit: SubmitHandler<IFormInput> = (data) => console.log(data);
  const [seePassword, setSeePassword] = useState<boolean>(false);

  const setSeePasswordMode = () => setSeePassword(!seePassword);

  return (
    <>
      <h2 className="text-2xl text-center font-bold mb-6">Login</h2>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="username">
            Username / Email
          </label>
          <input className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="username" type="text" placeholder="Enter your username / email"
            {...register("username", {
              required: {
                value: true,
                message: "Username/Email is required for login",
              },
            })}
          />
          <p className="text-red-300">{errors.username?.message}</p>
        </div>
        <div className="relative">
          <input
            className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline pr-10"
            id="password"
            type={seePassword ? "text" : "password"}
            placeholder="Enter your password"
            {...register("password", {
              required: { value: true, message: "Password is required" },
            })}
          />
          <button type="button" className="absolute top-0 right-0 mt-3 mr-3"
            onClick={setSeePasswordMode}>
            {seePassword ? (
              <FaEyeSlash className="text-gray-500" />
            ) : (
              <FaEye className="text-gray-500" />
            )}
          </button>
          <p className="text-red-300">{errors.password?.message}</p>
        </div>
        <div className="flex items-center justify-between mt-4">
          <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
            Login
          </button>
          <Link className="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800" href={ApplicationRoutesEnum.LOGIN}>
            Forgot Password?
          </Link>
        </div>
      </form>
    </>
  );
};

export default page;
