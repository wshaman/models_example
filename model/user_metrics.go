package model

func (m *modelMetrics) ListUsers() (ret []User, err error) {
    m.onStart("ListUsers")
    ret, err = m.delegate.ListUsers()
    m.onEnd("ListUsers")
    return ret, err
}
