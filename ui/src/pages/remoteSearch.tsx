import { useState } from 'react';

import { Release } from 'client/runic';

import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import OutboundRoundedIcon from '@mui/icons-material/OutboundRounded';
import Paper from '@mui/material/Paper';

import { ButtonMap, ButtonMapButton, LoadingIndicator } from '@dashotv/components';

import { ReleaseList, ReleasesForm, SearchForm, useSearchQuery } from 'components/releases';

const pagesize = 25;
const page = 1;
export interface RunicSearchProps {
  selector: (url: string) => void;
  selected?: string;
  rawForm: SearchForm;
}
const RemoteSearch = ({ rawForm, selector, selected }: RunicSearchProps) => {
  const [form, setForm] = useState<SearchForm>(() => {
    rawForm.verified = false;
    return rawForm;
  });
  const [defaultForm] = useState<SearchForm>(rawForm);

  const { data, isFetching } = useSearchQuery(pagesize, (page - 1) * pagesize, rawForm);

  const reset = () => {
    setForm(defaultForm);
  };

  const handleSelect = (row: Release) => {
    if (!row.download) {
      return;
    }
    selector(row.download);
  };

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
      <Paper elevation={1} sx={{ p: 2, mb: 2, width: '100%' }}>
        <ReleasesForm form={form} setForm={setForm} reset={reset} />
      </Paper>
      {isFetching ? <LoadingIndicator /> : null}
      {data?.Releases ? <ReleaseList data={data?.Releases} actions={actions} selected={selected} /> : null}
    </>
  );
};

export default RemoteSearch;
