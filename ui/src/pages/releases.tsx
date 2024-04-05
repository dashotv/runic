import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import { Indexer } from 'client';

import { MenuItem, Stack, TextField } from '@mui/material';

import { Option } from 'components/Form';
import { useIndexersAllQuery } from 'components/indexers';
import { ReleaseList, useReleasesSearchQuery } from 'components/releases';

const resolutions: Option[] = [
  { label: 'All', value: '' },
  { label: '720p', value: '720' },
  { label: '1080p', value: '1080' },
  { label: '2160p', value: '2160' },
];
const types: Option[] = [
  { label: 'All', value: '' },
  { label: 'Movie', value: 'movie' },
  { label: 'TV', value: 'tv' },
  { label: 'Anime', value: 'anime' },
];

const Releases = () => {
  const { page } = useParams();
  const [source, setSource] = useState('');
  const [kind, setKind] = useState('');
  const [resolution, setResolution] = useState('');
  const [group, setGroup] = useState('');
  const [website, setWebsite] = useState('');
  const [indexers, setIndexers] = useState<Option[]>([]);
  const { data: indexersData } = useIndexersAllQuery();

  const { data } = useReleasesSearchQuery(Number(page) || 1, 25, source, kind, resolution, group, website);

  useEffect(() => {
    setIndexers(() => {
      if (!indexersData?.result) return [];
      const list = [{ label: 'All', value: '' }];
      const indexers = indexersData.result
        .filter((indexer: Indexer) => indexer.name)
        .map((indexer: Indexer) => {
          return { label: indexer.name!, value: indexer.name! };
        });
      return list.concat(indexers);
    });
  }, [indexersData, setIndexers]);

  return (
    <>
      <Stack direction="row" spacing={2} alignItems="center" sx={{ mb: 2 }}>
        <TextField
          label="Source"
          select
          value={source}
          size="small"
          variant="standard"
          onChange={e => setSource(e.target.value)}
          sx={{ width: '125px' }}
        >
          {indexers.map((option: Option) => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        <TextField
          label="Type"
          select
          value={kind}
          size="small"
          variant="standard"
          onChange={e => setKind(e.target.value)}
          sx={{ width: '125px' }}
        >
          {types.map((option: Option) => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        <TextField
          label="Resolution"
          select
          value={resolution}
          size="small"
          variant="standard"
          onChange={e => setResolution(e.target.value)}
          sx={{ width: '125px' }}
        >
          {resolutions.map((option: Option) => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        <TextField
          label="Group"
          value={group}
          size="small"
          variant="standard"
          onChange={e => setGroup(e.target.value)}
        />
        <TextField
          label="Website"
          value={website}
          size="small"
          variant="standard"
          onChange={e => setWebsite(e.target.value)}
        />
      </Stack>

      {data && <ReleaseList data={data?.result} />}
    </>
  );
};

export default Releases;
