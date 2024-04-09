import { useCallback, useEffect, useState } from 'react';
import { Helmet } from 'react-helmet-async';
import { useSearchParams } from 'react-router-dom';
import { createSearchParams } from 'react-router-dom';

import { Box, Button } from '@mui/material';
import Pagination from '@mui/material/Pagination';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';

import { Container, LoadingIndicator } from '@dashotv/components';

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

export const encodeSearchParams = params => createSearchParams(params);
// TODO: useForm and @hookform/devtools, see: https://www.youtube.com/watch?v=sD9fZxMO1us
export default function Search() {
  const [searchParams, setSearchParams] = useSearchParams();
  const [form, setForm] = useState<SearchForm>(
    Object.assign(formDefaults, {
      text: searchParams.get('text') || '',
      type: searchParams.get('type') || '',
      resolution: searchParams.get('resolution') || '',
      exact: searchParams.get('exact') === 'true',
      verified: searchParams.get('verified') === 'true',
      uncensored: searchParams.get('uncensored') === 'true',
      bluray: searchParams.get('bluray') === 'true',
    }),
  );
  const [page, setPage] = useState(1);
  // const [qs, setQs] = useState(queryString(form));
  const { isFetching, data } = useSearchQuery(pagesize, (page - 1) * pagesize, searchParams.toString());

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
  }, [form]);

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
        <Paper sx={{ p: 1, mb: 2, width: '100%' }}>
          {isFetching && <LoadingIndicator />}
          <ReleasesForm form={form} setForm={setForm} reset={reset} />
        </Paper>
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
}
