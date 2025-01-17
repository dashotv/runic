import { useForm } from 'react-hook-form';

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
  indexers,
  setForm,
  reset,
}: {
  form: SearchForm;
  indexers?: Option[];
  setForm: (data: SearchForm) => void;
  reset?: () => void;
}) {
  const { handleSubmit, control } = useForm({ values: form });
  const submit = (data: SearchForm) => {
    setForm(data);
  };

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
          <Select name="source" control={control} options={indexers || []} />
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

export interface SearchEmbeddedForm {
  text: string;
  uncensored: boolean;
  bluray: boolean;
  verified: boolean;
  exact: boolean;
}

export function ReleasesEmbeddedForm({
  form,
  // indexers,
  setForm,
  reset,
}: {
  form: SearchForm;
  // indexers?: Option[];
  setForm: React.Dispatch<React.SetStateAction<SearchForm>>;
  reset?: () => void;
}) {
  const formToText = (form: SearchForm) => {
    const transformed = {
      text: '',
      uncensored: form.uncensored,
      bluray: form.bluray,
      verified: form.verified,
      exact: form.exact,
    };
    const text: string[] = [];
    if (form.text) text.push(form.text);
    if (form.year) text.push(form.year);
    if (form.episode) text.push(`${form.season || 0}x${form.episode}`);
    if (form.group) text.push(`group:${form.group}`);
    if (form.website) text.push(`website:${form.website}`);
    if (form.resolution) text.push(`res:${form.resolution}`);
    if (form.source) text.push(`source:${form.source}`);
    if (form.type) text.push(`type:${form.type}`);
    transformed.text = text.join(' ');
    return transformed;
  };

  const textToForm = (data: SearchEmbeddedForm) => {
    const form: SearchForm = {
      text: '',
      year: '',
      season: '',
      episode: '',
      group: '',
      website: '',
      resolution: '',
      source: '',
      type: '',
      uncensored: data.uncensored,
      bluray: data.bluray,
      verified: data.verified,
      exact: data.exact,
    };
    const parts = data.text.split(' ');
    const plain: string[] = [];
    for (const part of parts) {
      if (part.startsWith('group:')) {
        form.group = part.replace('group:', '');
      } else if (part.startsWith('website:')) {
        form.website = part.replace('website:', '');
      } else if (part.startsWith('res:')) {
        form.resolution = part.replace('res:', '');
      } else if (part.startsWith('source:')) {
        form.source = part.replace('source:', '');
      } else if (part.startsWith('type:')) {
        form.type = part.replace('type:', '');
      } else if (part.match(/^\d{1,2}x\d{1,3}$/)) {
        const [season, episode] = part.split('x');
        form.season = season;
        form.episode = episode;
      } else if (part.match(/^\d{4}$/)) {
        form.year = part;
      } else {
        if (part !== '' && part !== ' ') plain.push(part);
      }
    }
    form.text = plain.join(' ');
    return form;
  };

  const { handleSubmit, control } = useForm<SearchEmbeddedForm>({ values: formToText(form) });
  const submit = (data: SearchEmbeddedForm) => {
    setForm(textToForm(data));
  };

  return (
    <>
      <Box component="form" noValidate autoComplete="off" onSubmit={handleSubmit(submit)}>
        <Stack direction={{ xs: 'column', sm: 'row' }} spacing={1}>
          <Text name="text" label="search" control={control} />
          <Stack sx={{ pt: 1, pl: 2 }} direction="row" spacing={0}>
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
            <Button type="submit">Go</Button>
            <Button
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
