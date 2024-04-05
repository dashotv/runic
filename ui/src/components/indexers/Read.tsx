import { useState } from 'react';

import { Indexer } from 'client';
import { NZB } from 'client/newznab';

import { DialogActions, Typography } from '@mui/material';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';

import { Megabytes, Row } from 'components/common';

import { useRunicReadQuery } from '.';

export const IndexersRead = ({ indexer, handleClose }: { indexer: Indexer; handleClose: () => void }) => {
  if (!indexer.name) throw new Error('Indexer name is required');
  const [open, setOpen] = useState(true);
  const [cats] = useState<number[]>(indexer.categories ?? [5000, 2000, 5070]);
  const { isFetching, data } = useRunicReadQuery(indexer.name, cats);
  const close = () => {
    setOpen(false);
    handleClose();
  };

  return (
    <Dialog open={open} onClose={() => close()} fullWidth={true} maxWidth="xl">
      <DialogTitle>
        <Stack direction="row" spacing={2} width="100%" alignItems="baseline" justifyContent="space-between">
          <Stack direction="row" spacing={2} width="100%" alignItems="baseline">
            <Typography fontWeight="bolder" noWrap color="primary">
              Read
            </Typography>
            <Typography fontWeight="bolder" noWrap color="primary.dark">
              {indexer.name}
            </Typography>
            <Typography noWrap color="gray">
              ({data?.result?.length})
            </Typography>
          </Stack>
          <Stack direction="row" spacing={2} width="100%" alignItems="baseline" justifyContent="end">
            <Typography noWrap color="gray">
              {cats.join(',')}
            </Typography>
          </Stack>
        </Stack>
      </DialogTitle>
      <DialogContent>
        <Paper elevation={0} sx={{ width: '100%', p: 1 }}>
          {isFetching && <Typography variant="body1">Loading...</Typography>}
          {data?.result?.map((nzb: NZB) => (
            <Row key={nzb.id || nzb.infohash}>
              <Stack direction="row" spacing={1} width="100%" alignItems="baseline" justifyContent="space-between">
                <Typography fontWeight="bolder" minWidth="300px" noWrap color="primary">
                  {nzb.title}
                </Typography>
                <Stack direction="row" spacing={1} minWidth="225px" alignItems="baseline" justifyContent="end">
                  <Typography fontWeight="bolder" minWidth="125px" noWrap color="gray">
                    {nzb.category.join(',')}
                  </Typography>
                  <Megabytes value={nzb.size} ord="bytes" />
                </Stack>
              </Stack>
            </Row>
          ))}
        </Paper>
      </DialogContent>
      <DialogActions>
        <Button variant="contained" onClick={() => close()}>
          Ok
        </Button>
      </DialogActions>
    </Dialog>
  );
};
