export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <h1 className="text-2xl font-semibold">Ingredienti</h1>
      <h3 className="mt-2 text-gray-700">Selezionaaaa le tue preferenze riguardo agli ingredienti.</h3>
      {children}
    </div>
  );
}
