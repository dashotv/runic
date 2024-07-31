import { useCallback, useState } from 'react';

import { Release } from 'client/runic';

import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import OutboundRoundedIcon from '@mui/icons-material/OutboundRounded';
import { Grid, Pagination, Paper, Stack, Typography } from '@mui/material';

import { ButtonMap, ButtonMapButton, LoadingIndicator } from '@dashotv/components';

import { useIndexersOptionsQuery } from 'components/indexers';
import { ReleaseList, ReleasesEmbeddedForm, SearchForm, useSearchQuery } from 'components/releases';

const pagesize = 25;
export interface RunicSearchProps {
  selector: (url: string, tags: string) => void;
  selected?: string;
  rawForm: SearchForm;
}
const RemoteSearch = ({ rawForm, selector, selected }: RunicSearchProps) => {
  const [defaultForm] = useState<SearchForm>(rawForm);
  const [page, setPage] = useState(1);
  const [form, setForm] = useState<SearchForm>(() => {
    return rawForm;
  });

  const { data: indexers } = useIndexersOptionsQuery();
  const { data, isFetching } = useSearchQuery(pagesize, (page - 1) * pagesize, form);

  const reset = () => {
    setForm(defaultForm);
  };

  const handleSelect = (row: Release) => {
    if (!row.download) {
      return;
    }

    const tags: string[] = [];
    if (row.group) {
      tags.push(row.group);
    }
    if (row.website) {
      tags.push(row.website);
    }
    if (row.resolution) {
      tags.push(row.resolution);
    }

    selector(row.download, tags.join(' '));
  };
  const handleChange = useCallback((_event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  }, []);

  const actions = row => {
    const buttons: ButtonMapButton[] = [
      {
        Icon: OutboundRoundedIcon,
        color: 'primary',
        click: () => console.log('click'),
        title: 'edit',
      },
      {
        Icon: CheckCircleIcon,
        color: 'primary',
        click: () => handleSelect(row),
        title: 're-process',
      },
    ];
    return <ButtonMap buttons={buttons} size="small" />;
  };

  return (
    <>
      <Paper elevation={1} sx={{ p: 1, mb: 2 }}>
        {isFetching ? <LoadingIndicator /> : null}
        <Grid container spacing={1}>
          <Grid item md={8} xs={12}>
            <ReleasesEmbeddedForm form={form} setForm={setForm} reset={reset} indexers={indexers} />
          </Grid>
          <Grid item md={4} xs={12} justifyContent="end">
            <Stack direction="row" spacing={0} justifyContent="end" width="100%" sx={{ pt: 2 }}>
              <Typography variant="caption" color="disabled" textAlign="right" sx={{ pt: 0.8 }}>
                {data?.Total || 0}
              </Typography>
              <Pagination
                boundaryCount={0}
                page={page}
                count={Math.ceil((data?.Total || 0) / pagesize)}
                onChange={handleChange}
                sx={{
                  '& > .MuiPagination-ul': {
                    justifyContent: 'end',
                  },
                }}
              />
            </Stack>
          </Grid>
        </Grid>
      </Paper>
      {data?.Releases ? <ReleaseList data={data?.Releases} actions={actions} selected={selected} /> : null}
    </>
  );
};

export default RemoteSearch;
