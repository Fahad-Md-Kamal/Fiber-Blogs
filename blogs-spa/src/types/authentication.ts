
export type User = {
  token?: string;
  username: string;
  email: string;
  isSuperuser: boolean;
  isActive: boolean;
};

export const serializeResponseDataToUser = (responseData: any): User => {
  return {
    token: responseData?.token || undefined,
    username: responseData?.username || "",
    email: responseData?.email || "",
    isSuperuser: !!responseData?.is_super_user,
    isActive: !!responseData?.is_active,
  };
};
