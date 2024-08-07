import { useState } from 'react';

import { Indexer } from 'client/runic';

import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import ClassIcon from '@mui/icons-material/Class';
import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import PlayLessonIcon from '@mui/icons-material/PlayLesson';
import QueueIcon from '@mui/icons-material/Queue';
import SyncIcon from '@mui/icons-material/Sync';
import IconButton from '@mui/material/IconButton';
import Link from '@mui/material/Link';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Unstable_Grid2';

import { Container, LoadingIndicator, Published, Row } from '@dashotv/components';

import {
  IndexerDialog,
  useIndexerCreateMutation,
  useIndexerDeleteMutation,
  useIndexerMutation,
  useIndexerRefreshMutation,
  useIndexerSettingMutation,
  useIndexersAllQuery,
} from '.';
import { IndexersParse } from './Parse';
import { IndexersRead } from './Read';

export const IndexersList = () => {
  const [selected, setSelected] = useState<Indexer | undefined>(undefined);
  const [reading, setReading] = useState<Indexer | undefined>(undefined);
  const [parsing, setParsing] = useState<Indexer | undefined>(undefined);
  const { isFetching, data } = useIndexersAllQuery();
  const indexer = useIndexerMutation();
  const setting = useIndexerSettingMutation();
  const create = useIndexerCreateMutation();
  const destroy = useIndexerDeleteMutation();
  const refresh = useIndexerRefreshMutation();

  const handleClose = (data?: Indexer) => {
    setSelected(undefined);

    if (data) {
      if (data.id === '') {
        create.mutate(data);
        return;
      }
      indexer.mutate(data);
    }
  };

  const view = (row: Indexer) => {
    setSelected(row);
  };
  const read = (row: Indexer) => {
    setReading(row);
  };
  const parse = (row: Indexer) => {
    setParsing(row);
  };

  const toggle = (id?: string, name?: string, value?: boolean) => {
    if (!id || !name) return;
    setting.mutate({ id, setting: { name: name, value: value! } });
  };

  const newIndexer = () => {
    setSelected({
      id: '',
      name: '',
      url: '',
      active: false,
      categories: [],
    });
  };
  const deleteIndexer = (id?: string) => {
    if (!id) return;
    destroy.mutate(id);
  };
  const handleRefresh = (id: string) => {
    refresh.mutate(id);
  };

  return (
    <>
      <Container>
        <Grid container spacing={2}>
          <Grid xs={12} md={6}>
            <Stack spacing={1} direction="row" justifyContent="start" alignItems="center">
              <Typography variant="h6" color="primary">
                Indexers
              </Typography>
              <IconButton aria-label="refresh" color="primary" onClick={() => newIndexer()}>
                <QueueIcon />
              </IconButton>
              <IconButton aria-label="refresh" color="primary" title="refresh all" onClick={() => handleRefresh('all')}>
                <SyncIcon />
              </IconButton>
            </Stack>
          </Grid>
          <Grid xs={12} md={6}></Grid>
        </Grid>
      </Container>
      <Container>
        {isFetching && <LoadingIndicator />}

        <Paper elevation={0} sx={{ m: 0, width: '100%' }}>
          {data?.result?.map((row: Indexer) => (
            <Row key={row.id}>
              <Stack direction={{ xs: 'column', md: 'row' }} spacing={1} alignItems="center">
                <Stack
                  direction="row"
                  spacing={1}
                  width="100%"
                  maxWidth={{ xs: '100%', md: '900px' }}
                  pr="3px"
                  alignItems="center"
                >
                  <IconButton size="small" onClick={() => toggle(row.id, 'active', !row.active)} title="active">
                    <CheckCircleIcon color={row.active ? 'secondary' : 'disabled'} fontSize="small" />
                  </IconButton>
                  <Link href="#" onClick={() => view(row)}>
                    <Typography fontWeight="bolder" color="primary">
                      {row.name}
                    </Typography>
                  </Link>
                  <Categories categories={row.categories} />
                </Stack>
                <Stack direction="row" spacing={0} alignItems="center" width="100%" justifyContent="end">
                  {row.count !== undefined && row.count > 0 && <Typography variant="body2">{row.count}</Typography>}
                  {row.processed_at && <Published date={row.processed_at} />}
                  <IconButton size="small" onClick={() => read(row)} title="Read">
                    <ClassIcon color="primary" fontSize="small" />
                  </IconButton>
                  <IconButton size="small" onClick={() => parse(row)} title="Parse">
                    <PlayLessonIcon color="primary" fontSize="small" />
                  </IconButton>
                  <IconButton
                    aria-label="refresh"
                    color="primary"
                    title="refresh"
                    size="small"
                    onClick={() => row.name && handleRefresh(row.name)}
                  >
                    <SyncIcon fontSize="small" />
                  </IconButton>
                  <IconButton size="small" onClick={() => deleteIndexer(row.id)} title="active">
                    <DeleteForeverIcon color="error" fontSize="small" />
                  </IconButton>
                </Stack>
              </Stack>
            </Row>
          ))}
          {selected && <IndexerDialog handleClose={handleClose} indexer={selected} />}
          {reading && <IndexersRead indexer={reading} handleClose={() => setReading(undefined)} />}
          {parsing && <IndexersParse indexer={parsing} handleClose={() => setParsing(undefined)} />}
        </Paper>
      </Container>
    </>
  );
};

const Categories = ({ categories }: { categories?: number[] }) => {
  if (!categories || categories.length === 0) return null;
  return (
    <Stack direction="row" spacing={0.75} alignItems="baseline">
      <Typography variant="body1" color="secondary" fontWeight="bold" noWrap>
        categories
      </Typography>
      <Typography variant="body2" color="secondary.dark" noWrap>
        {categories?.length}
      </Typography>
    </Stack>
  );
};
