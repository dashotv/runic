import { HiOutlineNewspaper } from 'react-icons/hi2';
import { SiUtorrent, SiYoutube } from 'react-icons/si';

import SvgIcon from '@mui/material/SvgIcon';

export const DownloaderIcon = ({ downloader }: { downloader: string }) => {
  const component = (downloader: string) => {
    switch (downloader) {
      case 'nzb':
        return HiOutlineNewspaper;
      case 'torrent':
        return SiUtorrent;
      case 'metube':
        return SiYoutube;
      default:
        return SiUtorrent;
    }
  };

  return (
    <SvgIcon
      sx={{ width: '20px', height: '20px' }}
      component={component(downloader)}
      inheritViewBox
      fontSize="small"
      color="disabled"
    />
  );
};
