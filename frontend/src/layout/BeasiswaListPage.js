import { useEffect, useState } from 'react';
import BeasiswaDetail from '../components/BeasiswaDetail';
import BeasiswaList from '../components/BeasiswaList';
import { getBeasiswaList } from '../lib/beasiswa';

function BeasiswaListPage() {
  const [beasiswaList, setBeasiswaList] = useState(null);
  const [activeBeasiswa, setActiveBeasiswa] = useState(null);

  useEffect(() => {
    getBeasiswaList().then((d) => {
      setBeasiswaList(d.data);
    });
  }, []);

  return (
    <div>
      {!beasiswaList ? (
        <p>Loading</p>
      ) : (
        <div className="mx-auto grid max-w-screen-lg grid-cols-[1fr_2fr] gap-4 py-16 px-4">
          <BeasiswaList
            beasiswaList={beasiswaList}
            activeBeasiswa={activeBeasiswa}
            setActiveBeasiswa={setActiveBeasiswa}
          />
          {beasiswaList && <BeasiswaDetail beasiswa={activeBeasiswa} />}
        </div>
      )}
    </div>
  );
}

export default BeasiswaListPage;
