import { useEffect, useState } from 'react';
import { getBeasiswaList } from '../lib/beasiswa';
import LandingPageCard from '../components/LandingPageCard';

function LandingPage() {
  const [beasiswaList, setBeasiswaList] = useState([]);

  useEffect(() => {
    getBeasiswaList().then((d) => {
      setBeasiswaList(d.data);
    });
  }, []);

  return (
    <div>
      <div className=" py-20 px-64 text-center text-gray-800">
        <h2 className="mt-0 mb-10 text-5xl font-bold">
          Pilih Beasiswa Yang Terbaik Untuk Masa Depanmu
        </h2>
        <blockquote className="mb-8 text-lg font-semibold">
          <div>
            <p className="">
              Learning is a result of listening, which in turn leads to even
              better listening and attentiveness to the other person. In other
              words, to learn from the child, we must have empathy, and empathy
              grows as we learn.
            </p>
          </div>
        </blockquote>
        <cite className="block text-gray-600">- Alice Miller</cite>
      </div>
      <div className="py-6 text-center text-gray-800">
        <h2 className="text-2xl font-bold">Beasiswa Terkini</h2>
      </div>
      <div className="flex flex-wrap justify-center pb-8">
        {beasiswaList.slice(0, 4).map((item) => (
          <LandingPageCard
            key={item.judulBeasiswa}
            title={item.judulBeasiswa}
            description={item.deskripsi}
          />
        ))}
      </div>
    </div>
  );
}

export default LandingPage;
