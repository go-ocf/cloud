import PropTypes from 'prop-types'
import { useIntl } from 'react-intl'

import { Button } from '@/components/button'
import { Modal } from '@/components/modal'
import { messages as t } from './confirm-modal-i18n'

const NOOP = () => {}

export const ConfirmModal = ({
  onConfirm,
  confirmButtonText,
  cancelButtonText,
  title,
  body,
  loading,
  closeOnConfirm,
  show,
  onClose,
  ...rest
}) => {
  const { formatMessage: _ } = useIntl()

  const renderFooter = () => {
    return (
      <div className="w-100 d-flex justify-content-end align-items-center">
        <Button variant="secondary" onClick={onClose} disabled={loading}>
          {cancelButtonText || _(t.cancel)}
        </Button>
        <Button
          variant="primary"
          onClick={() => onConfirm(onClose)}
          loading={loading}
          disabled={loading}
        >
          {confirmButtonText || _(t.confirm)}
        </Button>
      </div>
    )
  }

  const renderBody = () => body

  return (
    <Modal
      {...rest}
      show={show}
      onClose={!loading ? onClose : NOOP}
      title={title}
      renderBody={renderBody}
      renderFooter={renderFooter}
      closeButton={!loading}
    />
  )
}

ConfirmModal.propTypes = {
  onConfirm: PropTypes.func.isRequired,
  title: PropTypes.node.isRequired,
  body: PropTypes.node.isRequired,
  confirmButtonText: PropTypes.string,
  cancelButtonText: PropTypes.string,
  loading: PropTypes.bool,
}

ConfirmModal.defaultProps = {
  confirmButtonText: null,
  cancelButtonText: null,
  loading: false,
}
