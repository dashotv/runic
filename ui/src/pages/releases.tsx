import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';

import { Indexer } from 'client';

import { MenuItem, Stack, TextField } from '@mui/material';

import { Container, Option } from '@dashotv/components';

import { useIndexersAllQuery } from 'components/indexers';
import { ReleaseList, useReleasesSearchQuery } from 'components/releases';
import { ReleaseTypes, Resolutions } from 'types/constants';

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
    <Container>
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
          {ReleaseTypes.map((option: Option) => (
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
          {Resolutions.map((option: Option) => (
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
    </Container>
  );
};

export default Releases;
