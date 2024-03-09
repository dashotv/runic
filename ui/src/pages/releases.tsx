import { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';

import { MenuItem, Stack, TextField } from '@mui/material';

import { Option } from 'components/Form';
import { Indexer, useIndexersAllQuery } from 'components/indexers';
import { ReleaseList, useReleasesQuery } from 'components/releases';

const resolutions: Option[] = [
  { label: 'All', value: '' },
  { label: '720p', value: '720p' },
  { label: '1080p', value: '1080p' },
  { label: '2160p', value: '2160p' },
];
const types: Option[] = [
  { label: 'All', value: '' },
  { label: 'Movie', value: 'movie' },
  { label: 'TV', value: 'tv' },
  { label: 'Anime', value: 'anime' },
];

const Releases = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [source, setSource] = useState('');
  const [type, setType] = useState('');
  const [resolution, setResolution] = useState('');
  const [group, setGroup] = useState('');
  const [website, setWebsite] = useState('');
  const [indexers, setIndexers] = useState<Option[]>([]);

  const { data } = useReleasesQuery(25, 0, searchParams.toString());
  const { data: indexersData } = useIndexersAllQuery();

  useEffect(() => {
    setSearchParams({ source, type, resolution, group, website });
  }, [setSearchParams, source, type, resolution, group, website]);

  useEffect(() => {
    setIndexers(() => {
      if (!indexersData) return [];
      return indexersData.results?.map((indexer: Indexer) => {
        return { label: indexer.name, value: indexer.name };
      });
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
          value={type}
          size="small"
          variant="standard"
          onChange={e => setType(e.target.value)}
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

      {data && <ReleaseList data={data?.results} />}
    </>
  );
};

export default Releases;
