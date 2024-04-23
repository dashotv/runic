import { PopularMovie } from 'client/runic';

import { PopularList } from './list';

export const PopularMovies = ({ mount, data }: { mount: string; data: PopularMovie[] }) => {
  return (
    <PopularList
      mount={mount}
      data={data.map((p: PopularMovie) => {
        return { title: `${p.id?.title} (${p.id?.year})`, count: p.count };
      })}
      type="verified"
    />
  );
};
