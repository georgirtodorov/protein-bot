import { useState, FormEvent, ChangeEvent } from "react";

export default function AddProtein() {
  // ✅ Explicitly type state as string
  const [amount, setAmount] = useState<string>("");
  const [description, setDescription] = useState<string>("");

  async function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    try {
      const res = await fetch("/v1/add", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          amount: Number(amount),
          description,
        }),
      });

      if (!res.ok) throw new Error("Failed to add protein");

      alert("Protein added!");
      setAmount("");
      setDescription("");
    } catch (err) {
      if (err instanceof Error) {
        alert(err.message);
      } else {
        console.error(err);
        alert("An unexpected error occurred.");
      }
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
          onChange={(e: ChangeEvent<HTMLInputElement>) =>
            setAmount(e.target.value)
          }
          className="w-full border p-2 rounded-lg"
          required
        />
        <input
          type="text"
          placeholder="Description (optional)"
          value={description}
          onChange={(e: ChangeEvent<HTMLInputElement>) =>
            setDescription(e.target.value)
          }
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
