import React from 'react';

import AccessTimeIcon from '@mui/icons-material/AccessTime';

import type { Meta, StoryObj } from '@storybook/react';

import { Pill } from './Pill';

//👇 This default export determines where your story goes in the story list
const meta: Meta<typeof Pill> = {
  component: Pill,
};

export default meta;

type Story = StoryObj<typeof Pill>;

export const Default: Story = {
  args: {
    name: 'Name',
    value: 'Value',
    color: 'primary',
  },
  render: args => <Pill {...args} />,
};
export const Column: Story = {
  args: {
    name: 'Name',
    value: 'Value',
    color: 'primary',
    variant: 'column',
  },
  render: args => <Pill {...args} />,
};
export const Icon: Story = {
  args: {
    name: 'Name',
    value: 'Value',
    color: 'secondary',
    icon: <AccessTimeIcon fontSize="small" />,
  },
  render: args => <Pill {...args} />,
};
export const Resolution: Story = {
  args: {
    name: 'R',
    value: '1080',
    color: 'gray',
  },
  render: args => <Pill {...args} />,
};
