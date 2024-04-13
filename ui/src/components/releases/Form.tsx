import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';

import { Indexer } from 'client/runic';

import CircleOutlinedIcon from '@mui/icons-material/CircleOutlined';
import CrisisAlertIcon from '@mui/icons-material/CrisisAlert';
import SportsBarIcon from '@mui/icons-material/SportsBar';
import SportsBarOutlinedIcon from '@mui/icons-material/SportsBarOutlined';
import VerifiedIcon from '@mui/icons-material/Verified';
import VerifiedOutlinedIcon from '@mui/icons-material/VerifiedOutlined';
import VideocamIcon from '@mui/icons-material/Videocam';
import VideocamOutlinedIcon from '@mui/icons-material/VideocamOutlined';
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';

import { IconCheckbox, Option, Select, Text } from '@dashotv/components';

import { useIndexersAllQuery } from 'components/indexers';
import { ReleaseTypes, Resolutions } from 'types/constants';

export interface SearchForm {
  text: string;
  year: string;
  season: string;
  episode: string;
  group: string;
  website: string;
  resolution: string;
  source: string;
  type: string;
  uncensored: boolean;
  bluray: boolean;
  verified: boolean;
  exact: boolean;
}

export function ReleasesForm({
  form,
  setForm,
  reset,
}: {
  form: SearchForm;
  setForm: React.Dispatch<React.SetStateAction<SearchForm>>;
  reset?: () => void;
}) {
  const { handleSubmit, control } = useForm({ values: form });
  const [indexers, setIndexers] = useState<Option[]>([]);
  const submit = (data: SearchForm) => {
    setForm(data);
  };
  const { data: indexersData } = useIndexersAllQuery();

  useEffect(() => {
    setIndexers(() => {
      if (!indexersData?.result) return [];
      return indexersData.result
        .filter((indexer: Indexer) => indexer.name)
        .map((indexer: Indexer) => {
          return { label: indexer.name!, value: indexer.name! };
        });
    });
  }, [indexersData, setIndexers]);
  return (
    <>
      <Box component="form" noValidate autoComplete="off" onSubmit={handleSubmit(submit)}>
        <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
          <Text sx={{ minWidth: '125px' }} name="text" label="search" control={control} />
          <Text name="year" control={control} />
          <Text name="season" control={control} />
          <Text name="episode" control={control} />
          <Text name="group" control={control} />
          <Text name="website" control={control} />
          <Select name="resolution" label="Rez" control={control} options={Resolutions} />
          <Select name="source" control={control} options={indexers} />
          <Select name="type" control={control} options={ReleaseTypes} />
          <Stack sx={{ pt: 1, pl: 2 }} direction="row" spacing={1}>
            <IconCheckbox
              name="exact"
              sx={{ mr: 0 }}
              icon={<CircleOutlinedIcon />}
              checkedIcon={<CrisisAlertIcon />}
              control={control}
            />
            <IconCheckbox
              name="verified"
              sx={{ mr: 0 }}
              icon={<VerifiedOutlinedIcon />}
              checkedIcon={<VerifiedIcon />}
              control={control}
            />
            <IconCheckbox
              name="bluray"
              sx={{ mr: 0 }}
              icon={<VideocamOutlinedIcon />}
              checkedIcon={<VideocamIcon />}
              control={control}
            />
            <IconCheckbox
              name="uncensored"
              sx={{ mr: 0 }}
              icon={<SportsBarOutlinedIcon />}
              checkedIcon={<SportsBarIcon />}
              control={control}
            />
            <Button variant="contained" type="submit">
              Go
            </Button>
            <Button
              variant="contained"
              onClick={() => {
                reset && reset();
              }}
            >
              Reset
            </Button>
          </Stack>
        </Stack>
      </Box>
    </>
  );
}
