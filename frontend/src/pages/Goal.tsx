import { useEffect, useState } from "react";
import { ProteinGoalType } from "../types/ProteinGoalType";
import { GoalService } from "../services/goal";

export default function Goal() {
  const [currentGoal, setCurrentGoal] = useState<ProteinGoalType | null>(null);
  const [newGoal, setNewGoal] = useState("");
  const [history, setHistory] = useState<ProteinGoalType[]>([]);
  const [refreshKey, setRefreshKey] = useState(0);

  // Fetch current goal and history
  useEffect(() => {
    async function fetchGoal() {
      try {
        const goalData = await GoalService.getGoal();
        setCurrentGoal(goalData);

        // If you have a history endpoint
        const historyData = await GoalService.getHistory?.(); // optional chaining if implemented
        if (historyData) setHistory(historyData);
      } catch (err) {
        console.error("Failed to load goal:", err);
      }
    }
    fetchGoal();
  }, [refreshKey]);

  // Submit new goal
  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    try {
      const goalNumber = Number(newGoal);
      if (isNaN(goalNumber) || goalNumber <= 0) {
        alert("Please enter a valid positive number");
        return;
      }

      await GoalService.setGoal(goalNumber);

      // Clear input and refresh data
      setNewGoal("");
      setRefreshKey(prev => prev + 1);
    } catch (err) {
      alert((err as Error).message);
    }
  }

  return (
    <div className="max-w-md mx-auto bg-white p-6 rounded-2xl shadow-md">
      <h2 className="text-xl font-bold mb-4 text-center">Protein Goal</h2>

      {/* Current Goal */}
      {currentGoal ? (
        <div className="p-3 bg-gray-100 rounded-xl mb-4 text-center">
          <p>
            <strong>Current goal:</strong> {currentGoal.amount}g
          </p>
          <p>
            <strong>Set on:</strong> {new Date(currentGoal.created_at).toLocaleDateString()}
          </p>
        </div>
      ) : (
        <div className="p-3 bg-gray-100 rounded-xl mb-4 text-center">
          Loading current goal...
        </div>
      )}

      {/* Form to set new goal */}
      <form className="space-y-4" onSubmit={handleSubmit}>
        <input
          type="number"
          placeholder="New goal (g)"
          value={newGoal}
          onChange={(e) => setNewGoal(e.target.value)}
          className="w-full border p-2 rounded-lg"
          required
        />
        <button
          type="submit"
          className="w-full bg-blue-500 text-white p-2 rounded-lg hover:bg-blue-600"
        >
          Set Goal
        </button>
      </form>

      {/* Optional: Goal History */}
      {history.length > 0 && (
        <div className="mt-6">
          <h3 className="text-lg font-bold mb-2">Goal History</h3>
          <ul className="border rounded-lg divide-y">
            {history.map((g) => (
              <li key={g.id} className="p-2">
                <span>{g.amount}g</span>{" "}
                <span className="text-gray-500 text-sm">
                  ({new Date(g.created_at).toLocaleDateString()})
                </span>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}
