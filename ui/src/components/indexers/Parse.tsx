import { useState } from 'react';

import { DialogActions, Typography } from '@mui/material';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import Stack from '@mui/material/Stack';

import { Indexer, IndexersReleases, useRunicParseQuery } from '.';

export const IndexersParse = ({ indexer, handleClose }: { indexer: Indexer; handleClose: () => void }) => {
  const [open, setOpen] = useState(true);
  const [cats] = useState<number[]>(indexer.categories ?? [5000, 2000, 5070]);
  const { isFetching, data } = useRunicParseQuery(indexer.name, cats);
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
              Parse
            </Typography>
            <Typography fontWeight="bolder" noWrap color="primary.dark">
              {indexer.name}
            </Typography>
            <Typography noWrap color="gray">
              ({data?.results?.length})
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
        {isFetching && <Typography variant="body1">Loading...</Typography>}
        {data?.results ? <IndexersReleases data={data?.results} /> : null}
      </DialogContent>
      <DialogActions>
        <Button variant="contained" onClick={() => close()}>
          Ok
        </Button>
      </DialogActions>
    </Dialog>
  );
};
