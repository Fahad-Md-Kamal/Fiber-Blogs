"use client";

import React, { useState } from "react";

import { useForm, SubmitHandler } from "react-hook-form";

interface IFormInput {
  username: String;
  email: String;
  password: String;
}

type Props = {};

const page = (props: Props) => {
  const { register, handleSubmit, formState } = useForm<IFormInput>();
  const {errors} = formState;
  const [valid, setValid] = useState<boolean>(false);
  const onSubmit: SubmitHandler<IFormInput> = (data) => console.log(data);

  return (
    <>
      <h2 className="text-2xl text-center font-bold mb-6">Signup</h2>
      
      <form onSubmit={handleSubmit(onSubmit)} noValidate>
        <div className="mb-4">
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="username"
          >
            Username
          </label>
          <input
            className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="username"
            type="text"
            placeholder="Enter your username"
            {...register("username", { required: {value:true, message: "username is required"}})}
          />
          <p className="text-red-300">{errors.username?.message}</p>
        </div>
        <div className="mb-4">
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="email"
          >
            Email
          </label>
          <input
            className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="email"
            type="email"
            placeholder="Enter your email address"
            {...register("email", 
            { 
                required: {value:true, message: "email required"},
                pattern: {
                    value:/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/,
                    message: "invalid email format"
                },
                validate: {
                    adminUser: (fieldValue) => {
                        return (
                            fieldValue !== "admin@gmail.com" || "Enter a different email address"
                        )
                    },
                    isBlackListed: (fieldValue) => {
                        return (
                            !fieldValue.endsWith("baddomain.com") || "This domain isn't supported."
                        )
                    }
                }
            })}
          />
          <p className="text-red-300">{errors.email?.message}</p>
        </div>
        <div className="mb-4">
          <label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="password"
          >
            Password
          </label>
          <input
            className="appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="password"
            type="password"
            placeholder="Enter your password"
            {...register("password", {required: {value:true, message: "password is required"}})}
          />
          <p className="text-red-300">{errors.password?.message}</p>
        </div>
        <div className="flex items-center justify-between">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Submit
          </button>
          <a
            className="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
            href="#"
          >
            Already has an account?
          </a>
        </div>
      </form>
    </>
  );
};

export default page;
