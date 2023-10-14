import type { LoaderFunction, MetaFunction } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import { GetAllUsers } from "~/api/users";

interface LoaderData {
  message: string;
  users: Awaited<ReturnType<typeof GetAllUsers>>;
}

export const loader: LoaderFunction = async () => {
  const users = await GetAllUsers()
  return { message: "Remix", users };
}

export default function Index() {
  const { message, users } = useLoaderData<LoaderData>();
  return (
    <main>
      <h1 className="text-4xl text-rose-500/80">Welcome to {message}!</h1>
      {users.map(user => (
        <div className="flex flex-col text-black" key={user.Id}>
          <li className="flex flex-col my-6">
            <Link to={`/users/${user.Id}`}>
              {user.FirstName} {user.LastName} {user.Email} 
            </Link>
          </li>
        </div>
      ))}
    </main>
  );
}

export const meta: MetaFunction = () => {
  return [
    { title: "oh, ok..." },
    { name: "description", content: "Welcome to the fray!" },
  ];
};
