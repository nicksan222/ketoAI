export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <div className="flex flex-row justify-between">
        <div>
          <h1 className="text-2xl font-semibold">In attesa</h1>
          <h3 className="mt-2 text-gray-700">
            Tutte le ricette in attesa di approvazione  aaa
          </h3>
        </div>
      </div>
      {children}
    </>
  );
}
