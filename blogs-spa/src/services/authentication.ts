// import { showFailureToaster, showSuccessToaster } from "../components/Tosters";
// import { User, serializeResponseDataToUser } from "../types/authentication";

import { User } from "../types/authentication";

// const PostApi = async <T>(urlPath: string, data: T, token?: string): Promise<T> => {
//   try {
//     const response = await fetch(`http://localhost:3000/${urlPath}`, {method: "POST",
//       headers: { "Content-Type": "application/json", Authorization: `Bearer ${token}`},
//       body: JSON.stringify(data),
//     });
//     let responseData = await response.json();
//     if (!response.ok) {
//       showFailureToaster(responseData.error);
//       return responseData;
//     }
//     return responseData;
//   } catch (error) {
//     showFailureToaster("Failed to login");
//   }
//   return {} as T;
// };

export const loginApiService = async <T>(data: unknown): Promise<User> => {
  const response = await fetch(`http://localhost:3000/login`, {
    headers: { "Content-Type": "application/json"},
    method: "POST",
    body: JSON.stringify(data),
  });

  const responseUser = (await response.json()) as User;
  return responseUser;
};
