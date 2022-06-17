import { useState } from 'react';
import LandingPageCard from '../components/LandingPageCard';

function LandingPage() {
  const dataCard = [
    {
      title: 'Mitra 1',
      description:
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
      image:
        'https://dicoding-web-img.sgp1.cdn.digitaloceanspaces.com/original/commons/new-ui-logo.png',
    },

    {
      title: 'Mitra 2',
      description:
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
      image:
        'https://lpdp.kemenkeu.go.id/storage/setting/logo/logo_1630922364.png',
    },

    {
      title: 'Mitra 3',
      description:
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
      image:
        'https://idcamp.ioh.co.id/images/indosat-ooredoo-hutchison-logo.png',
    },

    {
      title: 'Mitra 4',
      description:
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
      image:
        'https://www.kemdikbud.go.id/main/addons/shared_addons/themes/november_theme/img/kemdikbud_64x64.png',
    },

    {
      title: 'Mitra 5',
      description:
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
      image:
        'https://www.kemdikbud.go.id/main/addons/shared_addons/themes/november_theme/img/kemdikbud_64x64.png',
    },
  ];

  const [data] = useState(dataCard);

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
        {data.slice(0, 4).map((item) => (
          <LandingPageCard
            key={item.title}
            title={item.title}
            description={item.description}
            image={item.image}
          />
        ))}
      </div>
    </div>
  );
}

export default LandingPage;
