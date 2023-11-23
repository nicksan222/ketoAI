export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <h1 className="text-2xl font-semibold">Nuova ricetta</h1>
      <h3 className="mt-2 text-gray-700">Crea una nuova ricetta</h3>
      {children}
    </div>
  );
}
