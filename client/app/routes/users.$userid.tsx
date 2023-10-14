import { type LoaderFunction, type MetaFunction } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { GetUserById } from "~/api/users";


interface LoaderData {
  user: Awaited<ReturnType<typeof GetUserById>>;
}

// export const loader: LoaderFunction = async <T extends { userid: string }>({ params }: { params: T }) => {
export const loader: LoaderFunction = async ({ params }) => {
  console.log({ params })
  const user = await GetUserById(params.userid!)

  return { user }
}

export default function UserPage() {
  const { user } = useLoaderData<LoaderData>();

  return (
    <main>
      <h1>user  page</h1>
      <div>
        {JSON.stringify(user, null, 2)}
      </div>
    </main>
  )

}

export const meta: MetaFunction = () => {
  return [
    { title: "oh, ok..." },
    { name: "description", content: "Welcome to the BananaLand!" },
  ];
};
