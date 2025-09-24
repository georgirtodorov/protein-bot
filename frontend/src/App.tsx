import { BrowserRouter as Router, Route, Routes, Link, Navigate } from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import AddProtein from "./pages/AddProtein";
import History from "./pages/History";
import Goal from "./pages/Goal";

export default function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-50 text-gray-900">
        {/* Header */}
        <header className="bg-white shadow-md p-4 flex justify-between items-center">
          <h1 className="text-xl font-bold">Protein Bot</h1>
          <nav className="flex gap-4">
            <Link to="/dashboard">Dashboard</Link>
            <Link to="/add">Add</Link>
            <Link to="/history">History</Link>
            <Link to="/goal">Goal</Link>
          </nav>
        </header>

        {/* Pages */}
        <main className="p-6">
          <Routes>
            {/* Redirect root "/" to "/add" */}
            <Route path="/" element={<Navigate to="/add" replace />} />

            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/add" element={<AddProtein />} />
            <Route path="/history" element={<History />} />
            <Route path="/goal" element={<Goal />} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}
