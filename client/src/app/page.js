import Image from "next/image";

export default function Home() {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <div>
        <div>
          <div>create</div>
          <div>search</div>
        </div>
        <table class="table-auto">
          <thead>
            <tr>
              <th>name</th>
              <th>price</th>
              <th>status</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Chicken Pork with Radish</td>
              <td>30000</td>
              <td>available</td>
            </tr>
            <tr>
              <td>Chicken Pork</td>
              <td>35000</td>
              <td>not available</td>
            </tr>
          </tbody>
        </table>
      </div>

      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start"></main>
    </div>
  );
}
