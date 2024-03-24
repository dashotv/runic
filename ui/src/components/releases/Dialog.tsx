import { Link } from 'react-router-dom';

import CloseIcon from '@mui/icons-material/Close';
import Box from '@mui/material/Box';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import IconButton from '@mui/material/IconButton';
import Paper from '@mui/material/Paper';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import { useTheme } from '@mui/material/styles';
import useMediaQuery from '@mui/material/useMediaQuery';

import { Chrono, Group, Megabytes, Resolution } from 'components/common';

import { DownloaderIcon, Release } from '.';

export type ReleaseDialogProps = {
  open: boolean;
  handleClose: () => void;
  release: Release;
  // actions?: (row: Release) => JSX.Element;
};
export const ReleaseDialog = ({
  open,
  handleClose,
  // actions,
  release: {
    id,
    title,
    raw,
    description,
    published_at,
    source,
    type,
    group,
    website,
    resolution,
    size,
    verified,
    view,
    download,
    infohash,
    checksum,
    downloader,
    created_at,
    updated_at,
  },
}: ReleaseDialogProps) => {
  const theme = useTheme();
  const fullScreen = useMediaQuery(theme.breakpoints.down('md'));
  return (
    <Dialog open={open} onClose={handleClose} fullWidth fullScreen={fullScreen} maxWidth="md">
      <DialogTitle>
        <Typography noWrap color="primary">
          {title}
        </Typography>
        <Stack direction={{ xs: 'column', md: 'row' }} spacing={1} alignItems="center">
          <Stack width={{ xs: '100%', md: 'inherit' }} direction="row" spacing={1} alignItems="center">
            <DownloaderIcon downloader={downloader} verified={verified} />
            <Stack direction="row" spacing={1} alignItems="center">
              <Resolution resolution={resolution} variant="default" />
              <Group group={group} author={website} variant="default" />
            </Stack>
          </Stack>
          <Stack width={{ xs: '100%', md: 'inherit' }} direction="row" spacing={1} alignItems="center">
            <Typography variant="subtitle2" color="textSecondary" pl="3px">
              {source}:{type}
            </Typography>
            <Typography variant="subtitle2" color="textSecondary" pl="3px">
              {size ? `${size}mb` : ''}
            </Typography>
            <Typography variant="subtitle2" color="gray" pl="3px">
              {published_at && <Chrono fromNow>{published_at}</Chrono>}
            </Typography>
          </Stack>
        </Stack>
        <IconButton
          aria-label="close"
          onClick={handleClose}
          sx={{
            position: 'absolute',
            right: 8,
            top: 8,
            color: theme => theme.palette.grey[500],
          }}
        >
          <CloseIcon />
        </IconButton>
      </DialogTitle>
      <DialogContent>
        <Stack width="100%" direction="column" spacing={1}>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            ID
          </Typography>
          <Typography minHeight="24px" variant="body1" color="primary">
            {id}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Title
          </Typography>
          <Typography minHeight="24px" variant="body1" color="primary">
            {title}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Raw
          </Typography>
          <Typography minHeight="24px" variant="body1" color="primary">
            {raw?.title}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Size
          </Typography>
          <Typography component="div" minHeight="24px" variant="body1" color="primary">
            {size ? <Megabytes ord="bytes" value={size} /> : '?'}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Links
          </Typography>
          <div>
            {view && (
              <Link style={{ color: theme.palette.primary.main }} to={view}>
                <Typography minHeight="24px" variant="body1" color="primary">
                  View
                </Typography>
              </Link>
            )}
          </div>
          <div>
            {download && (
              <Link style={{ color: theme.palette.primary.main }} to={download}>
                <Typography minHeight="24px" variant="body1" color="primary">
                  Download
                </Typography>
              </Link>
            )}
          </div>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Hash
          </Typography>
          <Typography minHeight="24px" variant="body1" color="primary">
            {infohash}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Checksum
          </Typography>
          <Typography minHeight="24px" variant="body1" color="primary">
            {checksum}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Created
          </Typography>
          <Typography minHeight="24px" variant="body1" color="primary">
            {created_at}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Updated
          </Typography>
          <Typography minHeight="24px" variant="body1" color="primary">
            {updated_at}
          </Typography>
          <Typography variant="subtitle2" color="textSecondary" sx={{ position: 'relative', bottom: '-4px' }}>
            Description
          </Typography>
          <Paper elevation={0} sx={{ p: 2 }}>
            <Box>{description && <div dangerouslySetInnerHTML={{ __html: description }} />}</Box>
          </Paper>
        </Stack>
      </DialogContent>
      <DialogActions>{/* <Box>{actions && actions(release)}</Box> */}</DialogActions>
    </Dialog>
  );
};
