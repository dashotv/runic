import { useCallback, useEffect, useState } from 'react';
import { Helmet } from 'react-helmet-async';
import { createSearchParams, useSearchParams } from 'react-router-dom';

import { Box, Button, Pagination, Paper, Stack } from '@mui/material';

import { Container, LoadingIndicator } from '@dashotv/components';

import { useIndexersOptionsQuery } from 'components/indexers';
import { ReleaseList, ReleasesForm, SearchForm, useSearchQuery } from 'components/releases';

const pagesize = 25;
const formDefaults: SearchForm = {
  text: '',
  year: '',
  season: '',
  episode: '',
  group: '',
  website: '',
  resolution: '',
  source: '',
  type: '',
  uncensored: false,
  bluray: false,
  verified: false,
  exact: false,
};

const Search = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const [page, setPage] = useState(1);
  const [form, setForm] = useState<SearchForm>(
    Object.assign(formDefaults, {
      text: searchParams.get('text') || '',
      type: searchParams.get('type') || '',
      group: searchParams.get('group') || '',
      website: searchParams.get('author') || '',
      resolution: searchParams.get('resolution') || '',
      exact: searchParams.get('exact') === 'true',
      verified: searchParams.get('verified') === 'true',
      uncensored: searchParams.get('uncensored') === 'true',
      bluray: searchParams.get('bluray') === 'true',
    }),
  );
  const { isFetching, data } = useSearchQuery(pagesize, (page - 1) * pagesize, form);
  const { data: indexers } = useIndexersOptionsQuery();
  const encodeSearchParams = params => createSearchParams(params);

  const handleChange = useCallback((_event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  }, []);

  const reset = () => {
    setForm(formDefaults);
    setPage(1);
  };

  // const click = useCallback(() => {
  //   console.log('click');
  // }, []);

  // useEffect(() => {
  //   setQs(queryString(form));
  // }, [form, queryString]);

  useEffect(() => {
    setSearchParams(encodeSearchParams(form));
  }, [form, setSearchParams]);

  const setFormRift = () => {
    setForm(() => ({ ...formDefaults, source: 'rift', type: 'anime' }));
    setPage(1);
  };

  return (
    <>
      <Helmet>
        <title>Releases - Search</title>
        <meta name="description" content="A React Boilerplate application homepage" />
      </Helmet>

      <Container>
        {isFetching && <LoadingIndicator />}
        <Paper sx={{ p: 1, mb: 2, width: '100%' }}>
          <ReleasesForm form={form} setForm={setForm} reset={reset} indexers={indexers} />
        </Paper>
        {/* <Paper sx={{ p: 1, mb: 2, width: '100%' }}>
          <ReleasesEmbeddedForm form={form} setForm={setForm} reset={reset} indexers={indexers} />
        </Paper> */}
      </Container>

      <Container>
        <Paper sx={{ p: 1, mb: 2, width: '100%' }}>
          <Stack
            direction={{ xs: 'column', sm: 'row' }}
            spacing={1}
            alignItems="center"
            sx={{ width: '100%', justifyContent: 'space-between' }}
          >
            <Box>
              <Button variant="contained" onClick={setFormRift}>
                Rift
              </Button>
            </Box>
            {/* <ReleasesPresets {...{ setForm, setPage, formDefaults }} /> */}
            <Pagination
              boundaryCount={0}
              page={page}
              count={Math.ceil((data?.Total || 0) / pagesize)}
              onChange={handleChange}
            />
          </Stack>
        </Paper>
      </Container>

      <Container>{data?.Releases && <ReleaseList data={data.Releases} />}</Container>
    </>
  );
};

export default Search;
