import { useState } from "react";
import { ProteinService } from "../services/protein";

export default function AddProtein() {
  const [amount, setAmount] = useState("");
  const [description, setDescription] = useState("");

async function handleSubmit(e: React.FormEvent) {
  e.preventDefault();
  try {
    const data = await ProteinService.add({ amount: Number(amount), description });
    // Assuming apiRequest parses JSON and returns it
    alert(`Protein added: ${data.added}\nTotal: ${data.total}\nRemaining: ${data.remaining}`);
    setAmount("");
    setDescription("");
  } catch (err) {
    alert((err as Error).message);
  }
}

  return (
    <div className="max-w-md mx-auto bg-white p-6 rounded-2xl shadow-md">
      <h2 className="text-lg font-bold mb-4">Add Protein</h2>
      <form className="space-y-4" onSubmit={handleSubmit}>
        <input
          type="number"
          placeholder="Amount (g)"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          className="w-full border p-2 rounded-lg"
          required
        />
        <input
          type="text"
          placeholder="Description (optional)"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="w-full border p-2 rounded-lg"
        />
        <button
          type="submit"
          className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600"
        >
          Add
        </button>
      </form>
    </div>
  );
}