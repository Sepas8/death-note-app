import { Link } from "react-router";

const App = () => {
  return (
    <div className="flex flex-col items-center justify-center gap-4 h-screen">
      <Link to="/register" className="bg-red-500 text-white p-2 rounded-md">
        Register
      </Link>
    </div>
  );
};

export default App;
