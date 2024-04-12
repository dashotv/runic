import { ReactNode, useState } from 'react';
import { Link } from 'react-router-dom';
import Truncate from 'react-truncate-inside';

import { Release } from 'client';

import { Box, IconButton } from '@mui/material';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';

import { Chrono, Group, Megabytes, Resolution, Row } from '@dashotv/components';

import { DownloaderIcon } from '.';
import { ReleaseDialog } from './Dialog';
import { useReleaseSettingMutation } from './query';

export const ReleaseList = ({
  data,
  actions,
  selected,
}: {
  data: Release[];
  actions?: (row: Release) => ReactNode;
  selected?: string;
}) => {
  const [open, setOpen] = useState(false);
  const [viewing, setViewing] = useState<Release | null>(null);
  const releaseUpdate = useReleaseSettingMutation();
  const handleClose = () => setOpen(false);
  const view = (row: Release) => {
    setViewing(row);
    setOpen(true);
  };
  const toggle = (row: Release) => {
    if (!row.id) return;
    releaseUpdate.mutate(
      { id: row.id, setting: { name: 'verified', value: !row.verified } },
      {
        onSuccess: () => {
          row.verified = !row.verified;
        },
      },
    );
  };
  return (
    <>
      {data?.map((row: Release) => (
        <Row key={row.id} variant={selected == row.download ? 'selected' : undefined}>
          <Stack direction={{ xs: 'column', md: 'row' }} spacing={{ xs: 0, md: 1 }} alignItems="center">
            <Stack
              direction="row"
              spacing={1}
              width="100%"
              maxWidth={{ xs: '100%', md: '800px' }}
              pr="3px"
              alignItems="center"
            >
              <IconButton
                size="small"
                onClick={e => {
                  e.preventDefault();
                  toggle(row);
                }}
              >
                <DownloaderIcon downloader={row.downloader} verified={row.verified} />
              </IconButton>
              <Typography
                component="div"
                fontWeight="bolder"
                noWrap
                color="primary"
                sx={{ pr: 1, '& a': { color: 'primary.main' } }}
                title={row.raw?.title}
              >
                <Link
                  to="#"
                  onClick={e => {
                    e.preventDefault();
                    view(row);
                  }}
                >
                  <Truncate
                    text={
                      `${row.title}${row.year ? ` (${row.year})` : ''}${
                        row.episode ? ` [${row.season}x${row.episode}]` : ''
                      }` ||
                      row.raw?.title ||
                      ''
                    }
                    ellipsis=" ... "
                  />
                </Link>
              </Typography>
              <Stack
                display={{ xs: 'none', md: 'inherit' }}
                direction="row"
                spacing={1}
                alignItems="center"
                sx={{ pl: 1 }}
              >
                <Resolution resolution={row.resolution} variant="default" />
                <Group group={row.group} author={row.website} variant="default" />
                <Typography variant="subtitle2" noWrap color="textSecondary">
                  {row.source}:{row.type}
                </Typography>
              </Stack>
            </Stack>
            <Stack
              direction="row"
              spacing={1}
              alignItems="center"
              sx={{ width: '100%', justifyContent: { xs: 'start', md: 'end' } }}
            >
              <Stack
                display={{ xs: 'inherit', md: 'none' }}
                direction="row"
                spacing={1}
                alignItems="center"
                sx={{ pl: 1 }}
              >
                <Resolution resolution={row.resolution} variant="default" />
                <Group group={row.group} author={row.website} variant="default" />
                <Typography display={{ xs: 'none', md: 'inherit' }} variant="subtitle2" noWrap color="textSecondary">
                  {row.source}:{row.type}
                </Typography>
              </Stack>
              <Stack
                width={{ xs: '100%', md: 'auto' }}
                maxWidth="300px"
                direction="row"
                spacing={1}
                alignItems="center"
              >
                {/* <Typography fontWeight="bolder" minWidth="125px" noWrap color="gray">
                  {row.raw?.category?.join(',')}
                </Typography> */}
                {row.size ? <Megabytes ord="bytes" value={row.size} /> : null}
                <Typography noWrap variant="subtitle2" color="gray" pl="3px" width="100%">
                  {row.published_at && <Chrono fromNow>{row.published_at}</Chrono>}
                </Typography>
                <Box>{actions ? actions(row) : null}</Box>
              </Stack>
            </Stack>
          </Stack>
        </Row>
      ))}
      {viewing && <ReleaseDialog {...{ open, handleClose }} release={viewing} />}
    </>
  );
};
