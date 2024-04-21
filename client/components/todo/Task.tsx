"use client";
import { useEffect, useState } from "react";
import { Backy } from "@/components/ui/back";
import Footer from "@/components/ui/Footer";
import { TypewriterEffectSmooth } from "../ui/typo";
import { Todo } from "@/types/tasks";
import { todo } from "node:test";
import { get } from "node:http";
export function Mainbhai() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTask, setNewTask] = useState<string>("");

  const GetTodos = async () => {
    try {
      const response = await fetch("https://go-to-do.onrender.com/todo");
      const jsonData = await response.json();
      setTodos(jsonData);
    } catch (err) {
      console.error(err);
    }
  };
  const deleteTask = async (id: any) => {
    try {
      await fetch(`https://go-to-do.onrender.com/todo`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ id: id }),
      });
      setTodos(todos.filter((todo: Todo) => todo.id !== id));
    } catch (err) {
      console.error(err);
    }
  };
  const editTask = async (id: string, completed: boolean, title: string) => {
    try {
      await fetch(`https://go-to-do.onrender.com/todo/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ completed, title }),
      });
      setTodos(
        todos.map((todo: Todo) =>
          todo._id === id ? { ...todo, completed, title } : todo
        )
      );
    } catch (err) {
      console.error(err);
    }
  };

  const addTask = async (task: any) => {
    try {
      const response = await fetch("https://go-to-do.onrender.com/todo", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: newTask, completed: false }),
      });
      const jsonData = await response.json();
      GetTodos();
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    GetTodos();
  }, []);

  const words = [
    {
      text: "Just",
    },
    {
      text: "a",
    },
    {
      text: "simple",
    },
    {
      text: "todo",
    },
    {
      text: "list?",
      className:
        "text-blue-500 dark:text-blue-500  text-2xl md:text-6xl font-bold text-center",
    },
  ];
  return (
    <div className="w-full mx-auto rounded-md  h-screen overflow-hidden">
      <Backy
        backgroundColor="black"
        rangeY={800}
        particleCount={500}
        className="flex items-center flex-col -z-1 justify-center px-2 md:px-10 gap-7 py-4 w-full h-full"
      >
        <TypewriterEffectSmooth words={words} />

        <div className="flex flex-col sm:flex-row items-center gap-4 mt-6">
          <input
            type="text"
            value={newTask}
            onChange={(e) => setNewTask(e.target.value)}
            className="w-80 h-10 px-3 rounded-lg border-2 border-purple-500 bg-gray-900 text-purple-500 placeholder-purple-500 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent"
            placeholder="What's on your mind?"
          />
          <button
            onClick={addTask}
            className="px-4 py-2 bg-blue-600 hover:bg-blue-700 transition duration-200 rounded-lg text-white shadow-[0px_2px_0px_0px_#FFFFFF40_inset]"
          >
            Add something
          </button>
        </div>
        {todos.map((todo: any) => {
          return (
            <div key={todo.id} className="flex justify-center mt-4">
              <div className="bg-purple-400 bg-opacity-25 backdrop-blur-md p-4 rounded-lg shadow-lg">
                <h1 className="text-xl font-bold text-purple-800">
                  {todo.title}
                </h1>
                <p
                  className={todo.completed ? "text-green-500" : "text-red-500"}
                >
                  {todo.completed ? "Completed" : "Not Completed"}
                </p>
                <div className="flex justify-center mt-2">
                  <button
                    className="bg-purple-600 hover:bg-purple-700 text-white font-bold py-2 px-4 rounded mr-2"
                    onClick={() => deleteTask(todo.id)}
                  >
                    Delete
                  </button>
                  <button
                    className="bg-purple-600 hover:bg-purple-700 text-white font-bold py-2 px-4 rounded"
                    onClick={() =>
                      editTask(todo.id, todo.completed, todo.title)
                    }
                  >
                    Edit
                  </button>
                </div>
              </div>
            </div>
          );
        })}
      </Backy>

      <Footer />
    </div>
  );
}
