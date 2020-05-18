""" Function wrappers for logging and handling exceptions"""
from functools import wraps

from service.lib.utils import request_to_dict


def default_uncurried_logging_wrapper(func, log=None, default_return_value=None):
    """ Function decorator to wrap gRPC endpoints, log arguments, and handle/log exceptions"""
    @wraps(func)
    def wrap(*args, **kwargs):
        request_dict = request_to_dict(args[1])
        if log is not None:
            log.info("Called DeleteGeneralMerchant", extra=request_dict)
        return_value = default_return_value
        try:
            return_value = func(*args, **kwargs)
        except Exception as e:
            if log is not None:
                log.error("Error in call to DeleteGeneralMerchant: %s", repr(e), extra=request_dict, exc_info=e)
        return return_value
    return wrap
